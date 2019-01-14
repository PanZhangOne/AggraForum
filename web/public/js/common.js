var COMMON = {
    init() {
        var forms = document.getElementsByClassName('needs-validation');
        var validation = Array.prototype.filter.call(forms, function (form) {
            form.classList.add('was-validated');
        });

        window.AJAX = this.AJAX;
        window.MESSAGE = this.MESSAGE;
        window.UTIL = this.UTIL;
    },
    MESSAGE: {
        error: function (message) {
            var template = `<div class="alert alert-danger alert-dismissible fade show" role="alert">${message}<button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button></div>`;
            $('#message').append(template);
        },
        successTitleAndContent: function (title, message, href) {
            var link;
            if (href) link = `<a href="${href}" class="alert-link">点击跳转</a>`
            var template = `<div class="alert alert-success" role="alert">
                                <h4 class="alert-heading">${title}</h4>
                                <p>${message} ${link}</p>
                            </div>`;
            $('#message').append(template);
        }
    },
    AJAX: {
        post: function (submitBtn, formEl, successFn, errorFn) {
            if (!submitBtn || !formEl) return;
            var url = $(submitBtn).attr('data-url');
            var data = {};
            var inputs = $(formEl).find('input');
            var divContent = $(formEl).find('div[contenteditable]');
            var select = $(formEl).find('select');

            for (var i = 0; i < inputs.length; i++) {
                var isRequired = $(inputs[i]).attr('required');
                var key = $(inputs[i]).attr('name');
                var value = $(inputs[i]).val();
                if (isRequired && !value) return MESSAGE.error('请填写' + key);
                data[key] = value;
            }

            if (divContent || divContent.length) {
                for (var i = 0; i < divContent.length; i++) {
                    var key = $(divContent[i]).attr('name');
                    var value = $(divContent[i]).html();
                    data[key] = value;
                }
            }

            if (select || select.length) {
                for (var i = 0; i < select.length; i++) {
                    var key = $(select[i]).attr('name');
                    var value = $(select[i]).val();
                    data[key] = value;
                }
            }
            $.post({
                url,
                data,
                success: successFn,
                error: errorFn
            });
        }
    },
    UTIL: {
        firstUpperCase: function (str) {
            return str.toLowerCase().replace(/\b[a-z]/g, function (s) {
                return s.toUpperCase();
            });
        }
    }
};

COMMON.init();
