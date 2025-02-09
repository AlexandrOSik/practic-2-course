
(function() {

    let actionButton = document.getElementById('actionButton');

    if (actionButton) {
        actionButton.onclick = (e) => {
            fetch("", {method: "POST"}).then(e => {
                let outMessage = document.getElementById('outMessage');
                outMessage.innerHTML = '';
                outMessage.innerHTML = `<span class="appears">
                    Вы берёте флаг! Он кричит "Не для тебя моё знамя цвело!" и выпрыгивает из ваших рук. Нужно найти другой способ его забрать.
                </span>`;
            });
        };
    } else {
        setTimeout(() => { document.location = "/"; }, 20000);
    }

})();
