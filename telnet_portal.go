/*
 * Copyright (c) 2022 eightfivefour llc. All rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
 * WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package mud

import (
	"fmt"
	"github.com/mjolnir-engine/engine/uid"
	"github.com/rs/zerolog"
	"net"
)

type TelnetConfiguration struct {
	Host string
	Port int
}

type telnetConnection struct {
	id     uid.UID
	logger zerolog.Logger
	conn   net.Conn
}

func newTelnetConnection(p *telnetPortal, conn net.Conn) *telnetConnection {
	id := uid.New()

	return &telnetConnection{
		id:     id,
		logger: p.logger.With().Str("component", "telnet_connection").Str("id", (string)(id)).Logger(),
		conn:   conn,
	}
}

func (tc *telnetConnection) start() {
	tc.logger.Debug().Msg("starting connection")
	_, err := tc.conn.Write([]byte("Mjolnir MUD Engine\n"))

	if err != nil {
		tc.logger.Fatal().Err(err).Msg("failed to write to connection")
		tc.stop()
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := tc.conn.Read(buf)
			if err != nil {
				tc.logger.Fatal().Err(err).Msg("failed to read from connection")
				tc.stop()
				return
			}

			tc.logger.Debug().Msgf("read %d bytes", n)
		}
	}()
}

func (tc *telnetConnection) stop() {
	tc.logger.Info().Msg("stopping")
}

type telnetPortal struct {
	mud         *Mud
	config      *TelnetConfiguration
	logger      zerolog.Logger
	connections map[uid.UID]*telnetConnection
}

func newTelnetPortal(mud *Mud) *telnetPortal {
	return &telnetPortal{
		mud:    mud,
		config: mud.config.Telnet,
		logger: mud.Engine.Logger().With().
			Str("host", mud.config.Telnet.Host).
			Int("port", mud.config.Telnet.Port).
			Str("component", "telnet_portal").
			Logger(),
		connections: make(map[uid.UID]*telnetConnection),
	}
}

func (tp *telnetPortal) Start() {
	tp.logger.Info().Msg("starting")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", tp.config.Host, tp.config.Port))

	if err != nil {
		tp.logger.Fatal().Err(err).Msg("failed to start")
		panic(err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				tp.logger.Fatal().Err(err).Msg("failed to accept connection")
				panic(err)
			}

			tc := newTelnetConnection(tp, conn)
			tp.connections[tc.id] = tc

			tc.start()
		}
	}()
}

func (tp *telnetPortal) Stop() {
	tp.logger.Info().Msg("stopping")
}
