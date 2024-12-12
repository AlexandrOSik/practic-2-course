
document.forms[0].onsubmit = (e) => {
    e.preventDefault();
    let data = new FormData(e.target);
    fetch(e.target.getAttribute('action'), {
        body: data,
        method: "POST"
    }).then(evt => {
        if (evt.status != 200) {
            document.getElementById('outMessage').innerHTML = ``;
            document.getElementById('outMessage').innerHTML = `<span class="appears">Предмет ${item} успешно выскользнул из ваших рук. Обратитесь к Творцу за помощью.</span>`;
            return;
        }
        let item = document.getElementById('item').innerText;
        document.getElementById('actionMessage').innerHTML = `<span class="appears">Предмет ${item} успешно подобран</span>`;
        document.querySelectorAll("li").forEach(e => {
            if (e.innerText.startsWith(item)) {
                e.classList.toggle("own", true);
                e.classList.toggle("do-not-own", false);
            }
        });
    });
};




