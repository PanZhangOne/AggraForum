(function ($) {
    'use strict';
    var topic = {
        editors: {},
        eventBuild: function () {
            var self = this;
            // 点击回复按钮显示编辑框 START
            $('.reply-btn').on('click', function () {
                var replyID = $(this).attr('data-reply-id');
                if (!replyID) return;

                var replyItem = $("div[data-wrap-id$='" + replyID + "']");
                var replyEditorWrap = replyItem.find("div[data-id$='" + replyID + "']")[0];
                if (!replyEditorWrap) return;
                replyEditorWrap = $(replyEditorWrap);

                if (!replyEditorWrap.is(":hidden")) {
                    replyEditorWrap.fadeOut();
                    return;
                }
                var editorEl = replyEditorWrap.find('.editor-content')[0];
                if (!editorEl) return;
                editorEl = $(editorEl);
                editorEl.html("");
                editorEl.append(" <textarea id='" + replyID + "'></textarea>");
                var editor = new Editor();
                editor.render(document.getElementById(replyID));
                self.editors[replyID] = editor;
                replyEditorWrap.fadeIn();
            });
            // 点击回复按钮显示编辑框 END

            // 提交回复数据 START
            $('button.reply-submit').on('click', function () {
                var replyID = $(this).attr('data-reply-id');
                var topicID = $('#topic-id').attr('data-id');
                if (!replyID) return alert('数据出错，请刷新重试!');
                if (!topicID) return alert('数据出错，请刷新重试!');

                var editor = self.editors[replyID];
                var content = editor.codemirror.getValue();
                if (!content || content.length === 0)
                    return alert('请先填写回复内容!');
                var deviceInfo = UTIL.firstUpperCase(detector.os.name) + " " + detector.os.version + "/" +
                    UTIL.firstUpperCase(detector.browser.name) + " " +
                    detector.browser.version;

                $.ajax({
                    url: '/reply',
                    type: 'POST',
                    data: {parent_id: replyID, content: content, device_info: deviceInfo},
                    success: function (data) {
                        if (data.success) {
                            new dToast({body: '回复成功!'});
                            window.location.href = window.location.href;
                        } else {
                            new dToast({body: data.message})
                        }
                    },
                    error: function (err) {
                        new dToast(err.message || '网络错误');
                    }
                })
            });
            // 提交回复数据 END
        },
        init: function () {
            this.eventBuild();
        }
    };
    topic.init();
})(jQuery);
