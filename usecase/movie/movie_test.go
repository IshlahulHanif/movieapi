package movie

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/movieapi/pkg/httpclient"
	"reflect"
	"testing"
)

func TestModule_GetDetailByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		endpoint func() endpoint
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				endpoint: func() endpoint {
					mockHttpClient := httpclient.NewMockHttpClient(ctrl)
					mockHttpClient.EXPECT().DoRequest(gomock.Any(), gomock.Any()).Return(map[string]interface{}{}, nil)

					return endpoint{
						http: mockHttpClient,
					}
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "movie_id",
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "empty data",
			fields: fields{
				endpoint: func() endpoint {
					mockHttpClient := httpclient.NewMockHttpClient(ctrl)

					return endpoint{
						http: mockHttpClient,
					}
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "",
			},
			want:    map[string]interface{}{},
			wantErr: true,
		},
		{
			name: "error request",
			fields: fields{
				endpoint: func() endpoint {
					mockHttpClient := httpclient.NewMockHttpClient(ctrl)
					mockHttpClient.EXPECT().DoRequest(gomock.Any(), gomock.Any()).Return(map[string]interface{}{}, errors.New("err"))

					return endpoint{
						http: mockHttpClient,
					}
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "movie_id",
			},
			want:    map[string]interface{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Module{
				endpoint: tt.fields.endpoint(),
			}
			got, err := m.GetDetailByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDetailByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDetailByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_SearchByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		endpoint func() endpoint
	}
	type args struct {
		ctx       context.Context
		movieName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				endpoint: func() endpoint {
					mockHttpClient := httpclient.NewMockHttpClient(ctrl)
					mockHttpClient.EXPECT().DoRequest(gomock.Any(), gomock.Any()).Return(map[string]interface{}{}, nil)

					return endpoint{
						http: mockHttpClient,
					}
				},
			},
			args: args{
				ctx:       context.Background(),
				movieName: "movie name",
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "empty data",
			fields: fields{
				endpoint: func() endpoint {
					mockHttpClient := httpclient.NewMockHttpClient(ctrl)

					return endpoint{
						http: mockHttpClient,
					}
				},
			},
			args: args{
				ctx:       context.Background(),
				movieName: "",
			},
			want:    map[string]interface{}{},
			wantErr: true,
		},
		{
			name: "error request",
			fields: fields{
				endpoint: func() endpoint {
					mockHttpClient := httpclient.NewMockHttpClient(ctrl)
					mockHttpClient.EXPECT().DoRequest(gomock.Any(), gomock.Any()).Return(map[string]interface{}{}, errors.New("err"))

					return endpoint{
						http: mockHttpClient,
					}
				},
			},
			args: args{
				ctx:       context.Background(),
				movieName: "movie name",
			},
			want:    map[string]interface{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Module{
				endpoint: tt.fields.endpoint(),
			}
			got, err := m.SearchByName(tt.args.ctx, tt.args.movieName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
