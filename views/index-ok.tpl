{{ if .FlagCorrect }}
  <section>
    <div class="msg">
      <p>Вы проходите мимо огромного плоского бревна летящей походкой
      в перевалку с флагом <strong class="output-flag">{{ .Flag }}</strong> на плече!"</p>
    </div>
    <div class="msg">
      <p>Перед вами стоит <strong>кот Шрёдингера</strong> и показывает вам большой палец. 
        Вы, видимо, очень устали и не можете понять, куда он направлен.</p>
    </div>
    <div class="msg">
      <p>Нужно отдохнуть.</p>
    </div>
  </section>
{{ else }}
  <section>
    <div class="msg">
      <p>Вы проходите мимо огромного плоского бревна летящей походкой. 
        Через секунду вы летите летящим полётом в обратную сторону</p>
    </div>
    <div class="msg">
      <p>Было больно.</p>
    </div>
  </section>
  <section>
    <a href="/">Нужно возвращаться</a>
  </section>
{{ end }}
