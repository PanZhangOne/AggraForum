(function () {
    'use strict';

    const topic = {
        editors: {},
        methods: {
            show(el) {
                el.style.display = 'block';
            },
            hidden(el) {
                el.style.display = 'none';
            }
        },
        events: function () {
            const self = this;
            const showReplyBtns = document.querySelectorAll('.reply-btn');
            const replyBtns = document.querySelectorAll('button.reply-submit');
            const collectBtn = document.querySelector('#collect');
            const cancelCollectionBtn = document.querySelector('#cancel_collection');
            const likeBtn = document.querySelector('#like');
            const cancelLikeBtn = document.querySelector('#cancel_like');
            const cancelDislikeBtn = document.querySelector('#cancel_dislike');
            const dislikeBtn = document.querySelector('#dislike');
            const topicID = document.querySelector('#topic-id').getAttribute('data-id');

            // 显示回复框
            showReplyBtns.forEach(reply => {
                reply.addEventListener('click', function (event) {
                    const replyID = this.getAttribute('data-reply-id');
                    if (!replyID) return;

                    const replyItem = document.querySelector(`div[data-wrap-id="${replyID}"]`);
                    const replyWrap = replyItem.querySelector(`div[data-id="${replyID}"]`);
                    const replyWrapIsShow = replyWrap.style.display === 'block';
                    if (replyWrapIsShow) return self.methods.hidden(replyWrap);
                    self.methods.show(replyWrap);
                })
            });
            // 回复
            replyBtns.forEach(reply => {
                reply.addEventListener('click', function (event) {
                    const replyID = this.getAttribute('data-reply-id');
                    const content = document.getElementById(replyID).value;
                    const deviceInfo = window.deviceInfo;

                    if (!replyID) return alert('数据出错，请刷新重试!');
                    if (!topicID) return alert('数据出错，请刷新重试!');

                    axios.post('/reply', UTIL.stringify({
                        topic_id: topicID,
                        parent_id: replyID,
                        device_info: deviceInfo,
                        content
                    })).then(result => {
                        if (result.data.success) {
                            location.reload();
                        } else {
                            alert(result.data.message);
                        }
                    }).catch(() => alert('网络错误!'));
                })
            });

            // 收藏
            collectBtn && collectBtn.addEventListener('click', function (event) {
                axios.get(`/collect/topic/${topicID}`).then(result => {
                    alert(result.data.message);
                    if (result.data.success) location.reload();
                }).catch(err => alert(err.message));
            });

            // 取消收藏
            cancelCollectionBtn && cancelCollectionBtn.addEventListener('click', function (event) {
                axios.get(`/collect/topic/cancel/${topicID}`).then(result => {
                    alert(result.data.message);
                    if (result.data.success) location.reload();
                }).catch(err => alert(err.message));
            });

            // 喜欢
            likeBtn && likeBtn.addEventListener('click', function (event) {
                axios.get(`/topic/like/${topicID}`).then(result => {
                    alert(result.data.message);
                    if (result.data.success) location.reload();
                }).catch(err => alert(err.message));
            });
            // 取消喜欢
            cancelLikeBtn && cancelLikeBtn.addEventListener('click', function (event) {
                axios.get(`/topic/like/cancel/${topicID}`).then(result => {
                    alert(result.data.message);
                    if (result.data.success) location.reload();
                }).catch(err => alert(err.message));
            });

            // 不喜欢
            dislikeBtn && dislikeBtn.addEventListener('click', function (event) {
                axios.get(`/topic/dislike/${topicID}`).then(result => {
                    alert(result.data.message);
                    if (result.data.success) location.reload();
                }).catch(err => alert(err.message));
            });
            // 取消不喜欢
            cancelDislikeBtn && cancelDislikeBtn.addEventListener('click', function (event) {
                axios.get(`/topic/dislike/cancel/${topicID}`).then(result => {
                    alert(result.data.message);
                    if (result.data.success) location.reload();
                }).catch(err => alert(err.message));
            })

        }
    };

    topic.events();
})();
