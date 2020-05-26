package service

import (
	"context"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"xblog/dao"
	"xblog/model"
	pb "xblog/rpc"
	"xblog/utils"
)

func (s *Service) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (response *pb.CreatePostResponse, error error) {
	var post model.Post
	utils.StructCopy(request, &post)
	author, err := dao.GetUserById(post.AuthorId, "id,username")
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "the author not found")
	}
	post.AuthorId = author.ID
	post.Digest = utils.SubChineseString(post.Content, 0, 100)
	post.Status = 1
	post.ViewNumbers = 0
	postId, err := dao.CreatePost(post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	response = &pb.CreatePostResponse{PostId: uint32(postId)}
	return response, nil
}

func (s *Service) UpdatePost(ctx context.Context, request *pb.UpdatePostRequest) (response *pb.CommonResponse, error error) {
	var post model.Post
	utils.StructCopy(request, &post)
	logs.Info(post)
	post.Digest = utils.SubChineseString(post.Content, 0, 100)
	postId, err := dao.UpdatePost(post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CommonResponse{Result: int32(postId)}, nil
}

func (s *Service) PostList(ctx context.Context, request *pb.PostListRequest) (response *pb.PostListResponse, error error) {
	page := request.Page
	limit := request.Limit
	results, total, err := dao.GetPostList(limit, page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var postDetailArr []*pb.PostDetail
	for _, value := range results {
		postDetail := &pb.PostDetail{}
		postDetail.Id = uint32(value.ID)
		postDetail.Title = value.Title
		postDetail.ContentOrDigest = &pb.PostDetail_Content{Content: value.Content}
		//postDetail.TagArray
		postDetailArr = append(postDetailArr, postDetail)
	}

	response = &pb.PostListResponse{Count: total, Page: page, List: postDetailArr}
	return response, nil
}
