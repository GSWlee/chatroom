package dataorm

import (
	"log"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "a",
			args: args{data: User{
				Name:      "syg",
				Firstname: "s",
				Lastname:  "yg",
				Phone:     "110",
				Email:     "110",
				Password:  "123",
				Status:    "",
			}},
			wantErr: false,
		}, {
			name: "b",
			args: args{
				data: Room{
					Name:    "ccnu",
					Creator: "syg",
					Data:    "/ccnu.log",
				},
			},
			wantErr: false,
		}, {
			name: "c",
			args: args{
				data: Userinroom{
					Userid: 1,
					Roomid: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := insert(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_query(t *testing.T) {
	type args struct {
		tablename string
		cols      []string
		where     []string
		values    []string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "a",
			args: args{
				tablename: "User",
				cols:      nil,
				where:     nil,
				values:    nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "b",
			args: args{
				tablename: "User",
				cols:      []string{"i_d", "name"},
				where:     []string{"i_d", "name"},
				values:    []string{"1", "syg"},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := query(tt.args.tablename, tt.args.cols, tt.args.where, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println("Show")
			log.Println(got)
		})
	}
}

func Test_delete(t *testing.T) {
	type args struct {
		tablename string
		id        int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "a",
			args: args{
				tablename: "User",
				id:        1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := delete(tt.args.tablename, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_update(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "a",
			args: args{
				data: User{
					ID:       5,
					Phone:    "99999",
					Email:    "",
					Password: "",
					Status:   "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := update(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
