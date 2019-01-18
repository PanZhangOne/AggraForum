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
                    const topicID = document.querySelector('#topic-id').getAttribute('data-id');
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
            })
        }
    };

    topic.events();
})();
