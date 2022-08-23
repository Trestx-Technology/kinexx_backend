package router

import (
	"encoding/json"
	"kinexx_backend/pkg/entity"
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

var routes = entity.Routes{
	entity.Route{
		Name:        "Register",
		Method:      "POST",
		Pattern:     "/register",
		HandlerFunc: account_service.SignUp,
	},
	entity.Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/login",
		HandlerFunc: account_service.Login,
	},
	entity.Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/sociallogin",
		HandlerFunc: account_service.SocialMedialogin,
	},
	entity.Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/sendemailotp",
		HandlerFunc: account_service.ForgetPasswordOTPLink,
	},
	entity.Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/verifyotp",
		HandlerFunc: account_service.VerifyOTPAndSendToken,
	},
	entity.Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/changepassword",
		HandlerFunc: account_service.ChangePassword,
	},
	entity.Route{
		Name:        "update profile",
		Method:      "PUT",
		Pattern:     "/profile",
		HandlerFunc: profile_service.UpdateProfile,
	},
	entity.Route{
		Name:        "presignedurl",
		Method:      "POST",
		Pattern:     "/presignedurl",
		HandlerFunc: util_service.GetPreSignedURL,
	},
	entity.Route{
		Name:        "update profile",
		Method:      "PUT",
		Pattern:     "/password/profile",
		HandlerFunc: profile_service.ChangePassword,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "POST",
		Pattern:     "/profile",
		HandlerFunc: profile_service.SetProfile,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "POST",
		Pattern:     "/profile/update/movies",
		HandlerFunc: profile_service.UpdateMovies,
	},
	entity.Route{
		Name:        "Experience",
		Method:      "POST",
		Pattern:     "/experience",
		HandlerFunc: profile_service.AddExperience,
	},
	entity.Route{
		Name:        "Experience",
		Method:      "GET",
		Pattern:     "/experience/{userid}",
		HandlerFunc: profile_service.GetUserExperience,
	},
	entity.Route{
		Name:        "Experience",
		Method:      "PATCH",
		Pattern:     "/experience/{experienceId}",
		HandlerFunc: profile_service.UpdateExperience,
	},
	entity.Route{
		Name:        "block",
		Method:      "PATCH",
		Pattern:     "/block/{id}",
		HandlerFunc: profile_service.BlockUser,
	},
	entity.Route{
		Name:        "block",
		Method:      "PATCH",
		Pattern:     "/unblock/{id}",
		HandlerFunc: profile_service.BlockUser,
	},
	entity.Route{
		Name:        "Experience",
		Method:      "DELETE",
		Pattern:     "/experience/{experienceId}",
		HandlerFunc: profile_service.RemoveExperience,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "POST",
		Pattern:     "/profile/update/hobbies",
		HandlerFunc: profile_service.UpdateHobbies,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "POST",
		Pattern:     "/profile/select/portfoliovideo",
		HandlerFunc: profile_service.SelectPromoVideos,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "POST",
		Pattern:     "/profile/upload/portfoliovideo",
		HandlerFunc: profile_service.UpdatePromoVideos,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "POST",
		Pattern:     "/profile/upload/videos",
		HandlerFunc: profile_service.UpdateVideo,
	},
	entity.Route{
		Name:        "set profile",
		Method:      "PATCH",
		Pattern:     "/profile/changestatus",
		HandlerFunc: profile_service.ChangeStatus,
	},
	entity.Route{
		Name:        "get profile",
		Method:      "GET",
		Pattern:     "/profile",
		HandlerFunc: profile_service.Profile,
	},
	entity.Route{
		Name:        "get profile",
		Method:      "GET",
		Pattern:     "/profile/all",
		HandlerFunc: profile_service.GetAllProfile,
	},
	entity.Route{
		Name:        "get profile",
		Method:      "GET",
		Pattern:     "/profile/user/{userid}",
		HandlerFunc: profile_service.GetUserProfile,
	},
	entity.Route{
		Name:        "get profile",
		Method:      "GET",
		Pattern:     "/profile/search/{search}",
		HandlerFunc: profile_service.GetSearch,
	},
	entity.Route{
		Name:        "get profile",
		Method:      "GET",
		Pattern:     "/profile/post/{userID}",
		HandlerFunc: post.GetUserData,
	},
	entity.Route{
		Name:        "Posts",
		Method:      "POST",
		Pattern:     "/posts",
		HandlerFunc: post.AddPost,
	},
	entity.Route{
		Name:        "UpdatePosts",
		Method:      "PATCH",
		Pattern:     "/posts",
		HandlerFunc: post.UpdatePost,
	},
	entity.Route{
		Name:        "UpdatePosts",
		Method:      "DELETE",
		Pattern:     "/posts/{postID}",
		HandlerFunc: post.DeletePost,
	},
	entity.Route{
		Name:        "Posts",
		Method:      "GET",
		Pattern:     "/posts",
		HandlerFunc: post.GetPost,
	},
	entity.Route{
		Name:        "Posts",
		Method:      "GET",
		Pattern:     "/posts/{postID}",
		HandlerFunc: post.GetPostByPostID,
	},
	entity.Route{
		Name:        "Posts",
		Method:      "GET",
		Pattern:     "/stories",
		HandlerFunc: post.GetStories,
	},
	entity.Route{
		Name:        "Like Posts",
		Method:      "PATCH",
		Pattern:     "/like/posts/{postID}",
		HandlerFunc: post.LikePost,
	},
	entity.Route{
		Name:        "Like Posts",
		Method:      "GET",
		Pattern:     "/like/posts/{postID}",
		HandlerFunc: post.GetPostLikes,
	},
	entity.Route{
		Name:        "Share Posts",
		Method:      "GET",
		Pattern:     "/share/posts/{postID}",
		HandlerFunc: post.GetPostShares,
	},
	entity.Route{
		Name:        "Dislike Posts",
		Method:      "PATCH",
		Pattern:     "/dislike/posts/{postID}",
		HandlerFunc: post.DisLikePost,
	},
	entity.Route{
		Name:        "Posts",
		Method:      "GET",
		Pattern:     "/userposts",
		HandlerFunc: post.GetUserPost,
	},
	entity.Route{
		Name:        "Comments",
		Method:      "POST",
		Pattern:     "/comment",
		HandlerFunc: comment.AddComment,
	},
	entity.Route{
		Name:        "Dislike Comments",
		Method:      "PATCH",
		Pattern:     "/dislike/comments/{commentID}",
		HandlerFunc: comment.DisLikeComment,
	},
	entity.Route{
		Name:        "Like Comments",
		Method:      "PATCH",
		Pattern:     "/like/comments/{commentID}",
		HandlerFunc: comment.LikeComment,
	},
	entity.Route{
		Name:        "UpdateComments",
		Method:      "PATCH",
		Pattern:     "/comments",
		HandlerFunc: comment.UpdateComment,
	},
	entity.Route{
		Name:        "UpdateComment",
		Method:      "DELETE",
		Pattern:     "/comments/{commentID}",
		HandlerFunc: comment.DeleteComment,
	},
	entity.Route{
		Name:        "Comments",
		Method:      "GET",
		Pattern:     "/Comments/{postID}",
		HandlerFunc: comment.GetComment,
	},
	entity.Route{
		Name:        "Connections",
		Method:      "POST",
		Pattern:     "/connection",
		HandlerFunc: connection.AddConnection,
	},
	entity.Route{
		Name:        "Connections",
		Method:      "POST",
		Pattern:     "/updateconnection",
		HandlerFunc: connection.UpdateConnection,
	},
	entity.Route{
		Name:        "Connections",
		Method:      "GET",
		Pattern:     "/connection",
		HandlerFunc: connection.GetConnection,
	},
	entity.Route{
		Name:        "Connections",
		Method:      "GET",
		Pattern:     "/connection/online",
		HandlerFunc: connection.GetOnlineConnection,
	},
	entity.Route{
		Name:        "Connections",
		Method:      "GET",
		Pattern:     "/connection/online",
		HandlerFunc: connection.GetOnlineConnection,
	},
	entity.Route{
		Name:        "Connections",
		Method:      "GET",
		Pattern:     "/connection/count/{userID}",
		HandlerFunc: connection.GetConnectionCount,
	},
	entity.Route{
		Name:        "Share",
		Method:      "POST",
		Pattern:     "/share",
		HandlerFunc: share.AddShare,
	},
	entity.Route{
		Name:        "Share",
		Method:      "GET",
		Pattern:     "/share/{type}",
		HandlerFunc: share.GetShare,
	},
	entity.Route{
		Name:        "Share",
		Method:      "GET",
		Pattern:     "/share/my",
		HandlerFunc: share.GetMyShare,
	},
	entity.Route{
		Name:        "Chat",
		Method:      "POST",
		Pattern:     "/chat",
		HandlerFunc: chat.AddChat,
	},
	entity.Route{
		Name:        "Chat",
		Method:      "GET",
		Pattern:     "/chat/all",
		HandlerFunc: chat.GetAllChatListing,
	},
	entity.Route{
		Name:        "Chat",
		Method:      "GET",
		Pattern:     "/chat/{senderID}",
		HandlerFunc: chat.GetChat,
	},
	entity.Route{
		Name:        "Chat",
		Method:      "GET",
		Pattern:     "/spot/chat/{senderID}",
		HandlerFunc: chat.GetSpotChat,
	},
	entity.Route{
		Name:        "Chat",
		Method:      "DELETE",
		Pattern:     "/chat/{receiver}/{message}",
		HandlerFunc: chat.DeleteChat,
	},
	entity.Route{
		Name:        "Chat",
		Method:      "DELETE",
		Pattern:     "/hide-chat/{receiver}",
		HandlerFunc: chat.HideChat,
	},
	entity.Route{
		Name:        "Movie",
		Method:      "GET",
		Pattern:     "/movie/{query}/{page}",
		HandlerFunc: movies.FindMovies,
	},
	entity.Route{
		Name:        "Hobby",
		Method:      "GET",
		Pattern:     "/hobby/{query}/{page}",
		HandlerFunc: hobby.FindHobbies,
	},
	entity.Route{
		Name:        "Hobby",
		Method:      "POST",
		Pattern:     "/hobby",
		HandlerFunc: hobby.AddHobbies,
	},
	entity.Route{
		Name:        "Notification",
		Method:      "GET",
		Pattern:     "/notification",
		HandlerFunc: notification.FindNotifications,
	},
	entity.Route{
		Name:        "Notification",
		Method:      "PATCH",
		Pattern:     "/notification/{notification}",
		HandlerFunc: notification.UpdateNotifications,
	},
	entity.Route{
		Name:        "Rating",
		Method:      "POST",
		Pattern:     "/rating",
		HandlerFunc: rating_service.AddRating,
	},
	entity.Route{
		Name:        "Rating",
		Method:      "GET",
		Pattern:     "/rating/{userID}",
		HandlerFunc: rating_service.GetAllRatingByMe,
	},
	entity.Route{
		Name:        "Rating",
		Method:      "GET",
		Pattern:     "/rating/item/{itemID}",
		HandlerFunc: rating_service.GetItemRating,
	},
	entity.Route{
		Name:        "Rating",
		Method:      "GET",
		Pattern:     "/profile/with/rating/{userID}",
		HandlerFunc: rating_service.GetUserWithRating,
	},
	entity.Route{
		Name:        "Group",
		Method:      "POST",
		Pattern:     "/group",
		HandlerFunc: group.MakeGroup,
	},
	entity.Route{
		Name:        "Group",
		Method:      "GET",
		Pattern:     "/group",
		HandlerFunc: group.GetAllGroup,
	},
	entity.Route{
		Name:        "Group",
		Method:      "DELETE",
		Pattern:     "/group/{groupID}",
		HandlerFunc: group.DeleteGroup,
	},
	entity.Route{
		Name:        "Group",
		Method:      "PUT",
		Pattern:     "/group/{groupID}",
		HandlerFunc: group.EditGroup,
	},
	entity.Route{
		Name:        "Group",
		Method:      "GET",
		Pattern:     "/group/detail/{groupID}",
		HandlerFunc: group.GetGroupDetail,
	},
	entity.Route{
		Name:        "Group",
		Method:      "GET",
		Pattern:     "/group/user/{userID}",
		HandlerFunc: group.GetGroupByMe,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "POST",
		Pattern:     "/group/user/{groupID}/{userID}/{status}",
		HandlerFunc: groupuserservice.AddUserToGroup,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "POST",
		Pattern:     "/group/goal/{groupID}/{goalID}",
		HandlerFunc: goaluserservice.AddUserToGoal,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "DELETE",
		Pattern:     "/goal/{goalID}/{userID}",
		HandlerFunc: goaluserservice.RemoveUserFromGoal,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "GET",
		Pattern:     "/group/by/user/{userID}",
		HandlerFunc: groupuserservice.GetGroupsForUser,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "GET",
		Pattern:     "/users/in/group/{groupID}",
		HandlerFunc: groupuserservice.GetUsersInGroup,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "DELETE",
		Pattern:     "/group/{groupID}/{userID}",
		HandlerFunc: groupuserservice.RemoveUserFromGroup,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "GET",
		Pattern:     "/goal/in/group/{groupID}",
		HandlerFunc: goaluserservice.GetGoalsForUser,
	},
	entity.Route{
		Name:        "Group User",
		Method:      "GET",
		Pattern:     "/group/by/goal/{goalID}",
		HandlerFunc: goaluserservice.GetUsersInGoal,
	},
	entity.Route{
		Name:        "Goal",
		Method:      "POST",
		Pattern:     "/goal",
		HandlerFunc: goal.MakeGoal,
	},
	entity.Route{
		Name:        "Goal",
		Method:      "GET",
		Pattern:     "/goal",
		HandlerFunc: goal.GetAllGoals,
	},
	entity.Route{
		Name:        "Brand",
		Method:      "POST",
		Pattern:     "/brand",
		HandlerFunc: brand.Add,
	},
	entity.Route{
		Name:        "Brand",
		Method:      "GET",
		Pattern:     "/brand",
		HandlerFunc: brand.GetMany,
	},
	entity.Route{
		Name:        "Brand",
		Method:      "GET",
		Pattern:     "/search-brand/{keyword}",
		HandlerFunc: brand.Search,
	},
	entity.Route{
		Name:        "Brand",
		Method:      "GET",
		Pattern:     "/brand/{id}",
		HandlerFunc: brand.GetDetails,
	},
	entity.Route{
		Name:        "Subscription",
		Method:      "POST",
		Pattern:     "/subscription",
		HandlerFunc: subscription.Add,
	},
	entity.Route{
		Name:        "Subscription",
		Method:      "GET",
		Pattern:     "/subscription",
		HandlerFunc: subscription.GetMany,
	},
	entity.Route{
		Name:        "Subscription",
		Method:      "GET",
		Pattern:     "/subscription/{id}",
		HandlerFunc: subscription.GetDetails,
	},
	entity.Route{
		Name:        "Spot",
		Method:      "POST",
		Pattern:     "/spot",
		HandlerFunc: spot.Add,
	},
	entity.Route{
		Name:        "Spot",
		Method:      "GET",
		Pattern:     "/spot/{id}",
		HandlerFunc: spot.GetDetails,
	},
	entity.Route{
		Name:        "Spot",
		Method:      "DELETE",
		Pattern:     "/spot/{id}",
		HandlerFunc: spot.DeleteSpot,
	},
	entity.Route{
		Name:        "Spot",
		Method:      "GET",
		Pattern:     "/spot/{type}/{value}",
		HandlerFunc: spot.GetMany,
	},
	entity.Route{
		Name:        "Store",
		Method:      "POST",
		Pattern:     "/store",
		HandlerFunc: store.Add,
	},
	entity.Route{
		Name:        "Store",
		Method:      "GET",
		Pattern:     "/store/{id}",
		HandlerFunc: store.GetDetails,
	},
	entity.Route{
		Name:        "Store",
		Method:      "DELETE",
		Pattern:     "/store/{id}",
		HandlerFunc: store.DeleteStore,
	},
	entity.Route{
		Name:        "Store",
		Method:      "GET",
		Pattern:     "/store/{type}/{value}",
		HandlerFunc: store.GetMany,
	},
	entity.Route{
		Name:        "Product",
		Method:      "POST",
		Pattern:     "/product",
		HandlerFunc: product.Add,
	},
	entity.Route{
		Name:        "Product",
		Method:      "PUT",
		Pattern:     "/product/{id}",
		HandlerFunc: product.Update,
	},
	entity.Route{
		Name:        "Product",
		Method:      "GET",
		Pattern:     "/product/{id}",
		HandlerFunc: product.GetDetails,
	},
	entity.Route{
		Name:        "Product",
		Method:      "DELETE",
		Pattern:     "/product/{id}",
		HandlerFunc: product.DeleteProduct,
	},
	entity.Route{
		Name:        "Product",
		Method:      "GET",
		Pattern:     "/product/{type}/{value}",
		HandlerFunc: product.GetMany,
	},
	entity.Route{
		Name:        "HEALTH",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Health,
	},
}

func Health(w http.ResponseWriter, _ *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("social media login", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "health": "ok"})
	if err != nil {
		return
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login successfully", logrus.Fields{"duration": duration})
}
