package service

import (
	"context"
	"testing"

	sample_data "github.com/psinthorn/gostore/data_generator"
	"github.com/psinthorn/gostore/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()

	laptopWithNoId := sample_data.NewLaptop()
	laptopWithNoId.Id = ""

	laptopWithInvalidId := sample_data.NewLaptop()
	laptopWithInvalidId.Id = "invalid-id"

	laptopWithDulicateId := sample_data.NewLaptop()
	storeDuplicateId := NewInmemoryLaptopStore()
	err := storeDuplicateId.Save(laptopWithDulicateId)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		store  LaptopStore
		laptop *pb.Laptop
		code   codes.Code
	}{
		{
			name:   "create_success_with_id",
			laptop: sample_data.NewLaptop(),
			store:  NewInmemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "create_success_with_no_id",
			laptop: laptopWithNoId,
			store:  NewInmemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "create_failed_with_invalid_id",
			store:  NewInmemoryLaptopStore(),
			laptop: laptopWithInvalidId,
			code:   codes.InvalidArgument,
		},
		{
			name:   "create_failed_with_duplicate_id",
			store:  storeDuplicateId,
			laptop: laptopWithDulicateId,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
