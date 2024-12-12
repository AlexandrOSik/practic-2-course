<section>
  <div class="msg">Вы пришли {{ .Place }}. Перед вами появляется {{ .Teller }}.</div>
  <div class="msg">{{ .Teller }} вам говорит: {{ .Words }}.</div>
  <div class="msg">Вы немного озадачены.</div>
  {{ if .Form }}
    <div class="msg msg-item" id="actionMessage">
      <span>Вы что-то видите! Это - <strong id="item">{{ .Item }}</strong>.</span>
      <form action="{{.ActLink}}" method="post">
        <input type="hidden" name="check" value="{{.FormCheck}}">
        <input type="hidden" name="value" value="{{.FormValue}}">
        <button id="actionButton">Подобрать</button>
      </form>
    </div>
    <script src="/static/js/act.js"></script>
  {{ end }}
</section>

<p class="action-text">Но вы должны идти дальше. Выберите дорогу:</p>
<section><!--
  {{range $key, $val := .Links}}
  --><a href="{{ $val.Link }}">{{ $val.Text }}</a><!--
  {{end}}
--></section>
