const COMMON = (function () {

    const AJAX = (function () {
        function _isRequired(el) {
            return el.getAttribute('required');
        }

        function _getKeyAndValue(el) {
            const key = el.getAttribute('name');
            const value = el.value;
            return {key, value};
        }

        const post = function (submitBtn, formEl, successFn, errorFn) {
            if (!submitBtn || !formEl) return;
            const url = submitBtn.getAttribute('data-url');
            const _form = document.getElementById(formEl);
            const data = {};
            const inputs = _form.querySelectorAll('input');
            const divInputs = _form.querySelectorAll('div[contenteditable]');
            const selects = _form.querySelectorAll('select');

            inputs.forEach(input => {
                const {key, value} = _getKeyAndValue(input);
                data[key] = value;
            });
            divInputs.forEach(input => {
                const {key, value} = _getKeyAndValue(input);
                data[key] = value;
            });
            selects.forEach(input => {
                const {key, value} = _getKeyAndValue(input);
                data[key] = value;
            });
            axios.post(url, UTIL.stringify(data)).then(successFn).catch(errorFn);
        };

        return {post}
    })();

    const UTIL = (function () {
        function firstUpperCase(str) {
            return str.toLowerCase().replace(/\b[a-z]/g, function (s) {
                return s.toUpperCase();
            });
        }
        function stringify(data) {
            let res = '';
            for (const i in data) {
                res += `${i}=${data[i]}&`
            }
            return res;
        }

        return {firstUpperCase, stringify}
    })();

    const init = function () {
        const forms = document.getElementsByClassName('needs-validation');
        Array.prototype.filter.call(forms, function (form) {
            form.classList.add('was-validated');
        });

        window.deviceInfo = UTIL.firstUpperCase(detector.os.name) + " " + detector.os.version + "/" +
            UTIL.firstUpperCase(detector.browser.name) + " " +
            detector.browser.version;

        axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

        window.AJAX = AJAX;
        window.UTIL = UTIL;
    };

    return {init};
})();

COMMON.init();
