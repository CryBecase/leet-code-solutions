package main

import "container/list"

// https://leetcode-cn.com/problems/design-twitter/

type Twitter struct {
	userFollow map[int]map[int]int
	tweet      *list.List
}

type msg struct {
	userId  int
	tweetId int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		userFollow: make(map[int]map[int]int),
		tweet:      list.New(),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.userFollow[userId]; !ok {
		this.userFollow[userId] = make(map[int]int)
		this.userFollow[userId][userId] = 1
	}

	this.tweet.PushBack(msg{
		userId:  userId,
		tweetId: tweetId,
	})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	var follow map[int]int
	if _, ok := this.userFollow[userId]; !ok {
		return nil
	}
	follow = this.userFollow[userId]

	t := this.tweet.Back()
	result := make([]int, 0, 10)

	for t != nil && len(result) < 10 {
		id := t.Value.(msg).userId

		if _, ok := follow[id]; ok {
			// 此用户被关注了
			result = append(result, t.Value.(msg).tweetId)
		}

		t = t.Prev()
	}
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.userFollow[followerId]; !ok {
		this.userFollow[followerId] = make(map[int]int)
		this.userFollow[followerId][followerId] = 1
	}
	if _, ok := this.userFollow[followeeId]; !ok {
		this.userFollow[followeeId] = make(map[int]int)
		this.userFollow[followeeId][followeeId] = 1
	}

	this.userFollow[followerId][followeeId] = 1
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	if _, ok := this.userFollow[followerId]; !ok {
		this.userFollow[followerId] = make(map[int]int)
		this.userFollow[followerId][followerId] = 1
	}
	if _, ok := this.userFollow[followeeId]; !ok {
		this.userFollow[followeeId] = make(map[int]int)
		this.userFollow[followeeId][followeeId] = 1
	}

	delete(this.userFollow[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {

}
