package logger

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // os.setenv で環境変数を設定して、ログレベルを変更しているため、テストの並列実行はしない
func TestLogger_Debug(t *testing.T) {
	type fields struct {
		Logger func() (*Logger, *bytes.Buffer)
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "環境変数でログレベルを debug に変更したらDebugレベルのログは出力する",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "debug")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "debug message",
			},
			want: `{"time":"2024-01-23T12:34:56.000000789Z","level":"DEBUG","message":"debug message"}` + "\n",
		},
		{
			name: "ログレベルを warn 変更したら、Debugレベルのログは出力しない",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "warn")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "debug message",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, buf := tt.fields.Logger()
			logger.Debug(context.Background(), tt.args.msg, tt.args.fields...)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

//nolint:paralleltest // os.setenv で環境変数を設定して、ログレベルを変更しているため、テストの並列実行はしない
func TestLogger_Info(t *testing.T) {
	type fields struct {
		Logger func() (*Logger, *bytes.Buffer)
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "環境変数でログレベルを info に変更したらInfoレベルのログは出力する",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "info")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "info message",
			},
			want: `{"time":"2024-01-23T12:34:56.000000789Z","level":"INFO","message":"info message"}` + "\n",
		},
		{
			name: "ログレベルを warn 変更したら、Infoレベルのログは出力しない",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "warn")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "info message",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, buf := tt.fields.Logger()
			logger.Info(context.Background(), tt.args.msg, tt.args.fields...)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

//nolint:paralleltest // os.setenv で環境変数を設定して、ログレベルを変更しているため、テストの並列実行はしない
func TestLogger_Warn(t *testing.T) {
	type fields struct {
		Logger func() (*Logger, *bytes.Buffer)
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "環境変数でログレベルを warn に変更したらWarnレベルのログは出力する",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "warn")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "warn message",
			},
			want: `{"time":"2024-01-23T12:34:56.000000789Z","level":"WARN","message":"warn message"}` + "\n",
		},
		{
			name: "ログレベルを error 変更したら、Warnレベルのログは出力しない",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "error")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "warn message",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, buf := tt.fields.Logger()
			logger.Warn(context.Background(), tt.args.msg, tt.args.fields...)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

//nolint:paralleltest // os.setenv で環境変数を設定して、ログレベルを変更しているため、テストの並列実行はしない
func TestLogger_Error(t *testing.T) {
	type fields struct {
		Logger func() (*Logger, *bytes.Buffer)
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "ログレベルを error にしても、Errorレベルのログは出力される",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					clock := time.Date(2024, 1, 23, 12, 34, 56, 789, time.UTC)
					var buf bytes.Buffer
					_ = os.Setenv(LogLevelEnv, "error")
					return NewWithHandler(NewAppHandlerWithTime(&buf, clock)), &buf
				},
			},
			args: args{
				msg: "message",
				fields: []Field{
					FieldError(errors.New("error")),
				},
			},
			want: `{"time":"2024-01-23T12:34:56.000000789Z","level":"ERROR","message":"message","error":"error"}` + "\n",
		},
		{
			name: "ログレベルを fatal にしたら、Errorレベルのログは出力されない",
			fields: fields{
				Logger: func() (*Logger, *bytes.Buffer) {
					_ = os.Setenv(LogLevelEnv, "fatal")
					var buf bytes.Buffer
					return NewWithWriter(&buf), &buf
				},
			},
			args: args{
				msg: "message",
				fields: []Field{
					FieldError(errors.New("error")),
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, buf := tt.fields.Logger()
			logger.Error(context.Background(), tt.args.msg, tt.args.fields...)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
