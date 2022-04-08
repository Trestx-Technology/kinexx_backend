package router

import (
	"encoding/json"
	"kinexx_backend/pkg/services/account_service"
	chat "kinexx_backend/pkg/services/chat_service"
	comment "kinexx_backend/pkg/services/comment_service"
	connection "kinexx_backend/pkg/services/connection_service"
	hobby "kinexx_backend/pkg/services/hobby_service"
	movies "kinexx_backend/pkg/services/movies_service"
	notification "kinexx_backend/pkg/services/notification_service"
	post "kinexx_backend/pkg/services/post_service"
	"kinexx_backend/pkg/services/profile_service"
	"kinexx_backend/pkg/services/rating_service"
	share "kinexx_backend/pkg/services/share_service"
	"kinexx_backend/pkg/services/util_service"
	"time"

	"net/http"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"Register",
		"POST",
		"/register",
		account_service.SignUp,
	},
	Route{
		"Login",
		"POST",
		"/login",
		account_service.Login,
	},
	Route{
		"Login",
		"POST",
		"/sociallogin",
		account_service.SocialMedialogin,
	},
	Route{
		"Login",
		"POST",
		"/sendemailotp",
		account_service.ForgetPasswordOTPLink,
	},
	Route{
		"Login",
		"POST",
		"/verifyotp",
		account_service.VerifyOTPAndSendToken,
	},
	Route{
		"Login",
		"POST",
		"/changepassword",
		account_service.ChangePassword,
	},
	Route{
		"update profile",
		"PUT",
		"/profile",
		profile_service.UpdateProfile,
	},
	Route{
		"presignedurl",
		"POST",
		"/presignedurl",
		util_service.GetPreSignedURL,
	},
	Route{
		"update profile",
		"PUT",
		"/password/profile",
		profile_service.ChangePassword,
	},
	Route{
		"set profile",
		"POST",
		"/profile",
		profile_service.SetProfile,
	},
	Route{
		"set profile",
		"POST",
		"/profile/update/movies",
		profile_service.UpdateMovies,
	},
	Route{
		"set profile",
		"POST",
		"/profile/update/hobbies",
		profile_service.UpdateHobbies,
	},
	Route{
		"set profile",
		"POST",
		"/profile/select/portfoliovideo",
		profile_service.SelectPromoVideos,
	},
	Route{
		"set profile",
		"POST",
		"/profile/upload/portfoliovideo",
		profile_service.UpdatePromoVideos,
	},
	Route{
		"set profile",
		"POST",
		"/profile/upload/videos",
		profile_service.UpdateVideo,
	},
	Route{
		"set profile",
		"PATCH",
		"/profile/changestatus",
		profile_service.ChangeStatus,
	},
	Route{
		"get profile",
		"GET",
		"/profile",
		profile_service.Profile,
	},
	Route{
		"get profile",
		"GET",
		"/profile/all",
		profile_service.GetAllProfile,
	},
	Route{
		"get profile",
		"GET",
		"/profile/user/{userid}",
		profile_service.GetUserProfile,
	},
	Route{
		"get profile",
		"GET",
		"/profile/search/{search}",
		profile_service.GetSearch,
	},
	Route{
		"get profile",
		"GET",
		"/profile/post/{userID}",
		post.GetUserData,
	},
	Route{
		"Posts",
		"POST",
		"/posts",
		post.AddPost,
	},
	Route{
		"UpdatePosts",
		"PATCH",
		"/posts",
		post.UpdatePost,
	},
	Route{
		"UpdatePosts",
		"DELETE",
		"/posts/{postID}",
		post.DeletePost,
	},
	Route{
		"Posts",
		"GET",
		"/posts",
		post.GetPost,
	},
	Route{
		"Posts",
		"GET",
		"/posts/{postID}",
		post.GetPostByPostID,
	},
	Route{
		"Posts",
		"GET",
		"/stories",
		post.GetStories,
	},
	Route{
		"Like Posts",
		"PATCH",
		"/like/posts/{postID}",
		post.LikePost,
	},
	Route{
		"Like Posts",
		"GET",
		"/like/posts/{postID}",
		post.GetPostLikes,
	},
	Route{
		"Share Posts",
		"GET",
		"/share/posts/{postID}",
		post.GetPostShares,
	},
	Route{
		"Dislike Posts",
		"PATCH",
		"/dislike/posts/{postID}",
		post.DisLikePost,
	},
	Route{
		"Posts",
		"GET",
		"/userposts",
		post.GetUserPost,
	},
	Route{
		"Comments",
		"POST",
		"/comment",
		comment.AddComment,
	},
	Route{
		"Dislike Comments",
		"PATCH",
		"/dislike/comments/{commentID}",
		comment.DisLikeComment,
	},
	Route{
		"Like Comments",
		"PATCH",
		"/like/comments/{commentID}",
		comment.LikeComment,
	},
	Route{
		"UpdateComments",
		"PATCH",
		"/comments",
		comment.UpdateComment,
	},
	Route{
		"UpdateComment",
		"DELETE",
		"/comments/{commentID}",
		comment.DeleteComment,
	},
	Route{
		"Comments",
		"GET",
		"/Comments/{postID}",
		comment.GetComment,
	},
	Route{
		"Connections",
		"POST",
		"/connection",
		connection.AddConnection,
	},
	Route{
		"Connections",
		"POST",
		"/updateconnection",
		connection.UpdateConnection,
	},
	Route{
		"Connections",
		"GET",
		"/connection",
		connection.GetConnection,
	},
	Route{
		"Connections",
		"GET",
		"/connection/online",
		connection.GetOnlineConnection,
	},
	Route{
		"Connections",
		"GET",
		"/connection/online",
		connection.GetOnlineConnection,
	},
	Route{
		"Connections",
		"GET",
		"/connection/count/{userID}",
		connection.GetConnectionCount,
	},
	Route{
		"Share",
		"POST",
		"/share",
		share.AddShare,
	},
	Route{
		"Share",
		"GET",
		"/share/{type}",
		share.GetShare,
	},
	Route{
		"Share",
		"GET",
		"/share/my",
		share.GetMyShare,
	},
	Route{
		"Chat",
		"POST",
		"/chat",
		chat.AddChat,
	},
	Route{
		"Chat",
		"GET",
		"/chat/all",
		chat.GetAllChatListing,
	},
	Route{
		"Chat",
		"GET",
		"/chat/{senderID}",
		chat.GetChat,
	},
	Route{
		"Chat",
		"DELETE",
		"/chat/{receiver}/{message}",
		chat.DeleteChat,
	},
	Route{
		"Movie",
		"GET",
		"/movie/{query}/{page}",
		movies.FindMovies,
	},
	Route{
		"Hobby",
		"GET",
		"/hobby/{query}/{page}",
		hobby.FindHobbies,
	},
	Route{
		"Hobby",
		"POST",
		"/hobby",
		hobby.AddHobbies,
	},
	Route{
		"Notification",
		"GET",
		"/notification",
		notification.FindNotifications,
	},
	Route{
		"Notification",
		"PATCH",
		"/notification/{notification}",
		notification.UpdateNotifications,
	},
	Route{
		"Rating",
		"POST",
		"/rating",
		rating_service.AddRating,
	},
	Route{
		"Rating",
		"GET",
		"/rating/{userID}",
		rating_service.GetAllRatingByMe,
	},
	Route{
		"Rating",
		"GET",
		"/rating/item/{itemID}",
		rating_service.GetItemRating,
	},
	Route{
		"Rating",
		"GET",
		"/profile/with/rating/{userID}",
		rating_service.GetUserWithRating,
	},

	Route{
		"HEALTH",
		"GET",
		"/",
		Health,
	},
}

func Health(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("social media login", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "health": "ok"})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login successfully", logrus.Fields{"duration": duration})
}
