package router

import (
	"encoding/json"
	"kinexx_backend/pkg/services/account_service"
	brand "kinexx_backend/pkg/services/brand_service"
	chat "kinexx_backend/pkg/services/chat_service"
	comment "kinexx_backend/pkg/services/comment_service"
	connection "kinexx_backend/pkg/services/connection_service"
	goal "kinexx_backend/pkg/services/goal_service"
	goaluserservice "kinexx_backend/pkg/services/goal_user_service"
	group "kinexx_backend/pkg/services/group_service"
	groupuserservice "kinexx_backend/pkg/services/group_user_service"
	hobby "kinexx_backend/pkg/services/hobby_service"
	movies "kinexx_backend/pkg/services/movies_service"
	notification "kinexx_backend/pkg/services/notification_service"
	post "kinexx_backend/pkg/services/post_service"
	product "kinexx_backend/pkg/services/product_service"
	"kinexx_backend/pkg/services/profile_service"
	"kinexx_backend/pkg/services/rating_service"
	share "kinexx_backend/pkg/services/share_service"
	spot "kinexx_backend/pkg/services/spot_service"
	store "kinexx_backend/pkg/services/store_service"
	subscription "kinexx_backend/pkg/services/subscription_service"
	"kinexx_backend/pkg/services/util_service"
	"time"

	"net/http"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
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
		"Experience",
		"POST",
		"/experience",
		profile_service.AddExperience,
	},
	Route{
		"Experience",
		"GET",
		"/experience/{userid}",
		profile_service.GetUserExperience,
	},
	Route{
		"Experience",
		"PATCH",
		"/experience/{experienceId}",
		profile_service.UpdateExperience,
	},
	Route{
		"block",
		"PATCH",
		"/block/{id}",
		profile_service.BlockUser,
	},
	Route{
		"block",
		"PATCH",
		"/unblock/{id}",
		profile_service.BlockUser,
	},
	Route{
		"Experience",
		"DELETE",
		"/experience/{experienceId}",
		profile_service.RemoveExperience,
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
		"GET",
		"/spot/chat/{senderID}",
		chat.GetSpotChat,
	},
	Route{
		"Chat",
		"DELETE",
		"/chat/{receiver}/{message}",
		chat.DeleteChat,
	},
	Route{
		"Chat",
		"DELETE",
		"/hide-chat/{receiver}",
		chat.HideChat,
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
		"Group",
		"POST",
		"/group",
		group.MakeGroup,
	},
	Route{
		"Group",
		"GET",
		"/group",
		group.GetAllGroup,
	},
	Route{
		"Group",
		"DELETE",
		"/group/{groupID}",
		group.DeleteGroup,
	},
	Route{
		"Group",
		"PUT",
		"/group/{groupID}",
		group.EditGroup,
	},
	Route{
		"Group",
		"GET",
		"/group/detail/{groupID}",
		group.GetGroupDetail,
	},
	Route{
		"Group",
		"GET",
		"/group/user/{userID}",
		group.GetGroupByMe,
	},
	Route{
		"Group User",
		"POST",
		"/group/user/{groupID}/{userID}/{status}",
		groupuserservice.AddUserToGroup,
	},
	Route{
		"Group User",
		"POST",
		"/group/goal/{groupID}/{goalID}",
		goaluserservice.AddUserToGoal,
	},
	Route{
		"Group User",
		"DELETE",
		"/goal/{goalID}/{userID}",
		goaluserservice.RemoveUserFromGoal,
	},
	Route{
		"Group User",
		"GET",
		"/group/by/user/{userID}",
		groupuserservice.GetGroupsForUser,
	},
	Route{
		"Group User",
		"GET",
		"/users/in/group/{groupID}",
		groupuserservice.GetUsersInGroup,
	},
	Route{
		"Group User",
		"DELETE",
		"/group/{groupID}/{userID}",
		groupuserservice.RemoveUserFromGroup,
	},
	Route{
		"Group User",
		"GET",
		"/goal/in/group/{groupID}",
		goaluserservice.GetGoalsForUser,
	},
	Route{
		"Group User",
		"GET",
		"/group/by/goal/{goalID}",
		goaluserservice.GetUsersInGoal,
	},
	Route{
		"Goal",
		"POST",
		"/goal",
		goal.MakeGoal,
	},
	Route{
		"Goal",
		"GET",
		"/goal",
		goal.GetAllGoals,
	},
	Route{
		"Brand",
		"POST",
		"/brand",
		brand.Add,
	},
	Route{
		"Brand",
		"GET",
		"/brand",
		brand.GetMany,
	},
	Route{
		"Brand",
		"GET",
		"/search-brand/{keyword}",
		brand.Search,
	},
	Route{
		"Brand",
		"GET",
		"/brand/{id}",
		brand.GetDetails,
	},
	Route{
		"Subscription",
		"POST",
		"/subscription",
		subscription.Add,
	},
	Route{
		"Subscription",
		"GET",
		"/subscription",
		subscription.GetMany,
	},
	Route{
		"Subscription",
		"GET",
		"/subscription/{id}",
		subscription.GetDetails,
	},
	Route{
		"Spot",
		"POST",
		"/spot",
		spot.Add,
	},
	Route{
		"Spot",
		"GET",
		"/spot/{id}",
		spot.GetDetails,
	},
	Route{
		"Spot",
		"DELETE",
		"/spot/{id}",
		spot.DeleteSpot,
	},
	Route{
		"Spot",
		"GET",
		"/spot/{type}/{value}",
		spot.GetMany,
	},
	Route{
		"Store",
		"POST",
		"/store",
		store.Add,
	},
	Route{
		"Store",
		"GET",
		"/store/{id}",
		store.GetDetails,
	},
	Route{
		"Store",
		"DELETE",
		"/store/{id}",
		store.DeleteStore,
	},
	Route{
		"Store",
		"GET",
		"/store/{type}/{value}",
		store.GetMany,
	},
	Route{
		"Product",
		"POST",
		"/product",
		product.Add,
	},
	Route{
		"Product",
		"PUT",
		"/product/{id}",
		product.Update,
	},
	Route{
		"Product",
		"GET",
		"/product/{id}",
		product.GetDetails,
	},
	Route{
		"Product",
		"DELETE",
		"/product/{id}",
		product.DeleteProduct,
	},
	Route{
		"Product",
		"GET",
		"/product/{type}/{value}",
		product.GetMany,
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
