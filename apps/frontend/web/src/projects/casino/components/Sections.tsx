export function RegCurrency(){
return (
<section aria-labelledby="reg" className="py-6">
<div className="max-w-[1100px] mx-auto px-6 grid md:grid-cols-2 gap-4">
<div className="card p-6">
<h3 id="reg" className="mb-2 text-xl font-semibold">Регистрация</h3>
<p className="muted">Профиль с ником/аватаром/кошельком. После регистрации начисляется стартовый пакет монет <em>(настраивается)</em>.</p>
<ul className="grid gap-2 mt-2">
{[
'Email/пароль или OAuth (будущее).',
'Верификация почты; базовые лимиты против мульти‑аккаунтов.',
'Профиль: баланс, история транзакций, достижения.'
].map((x,i)=> (
<li key={i} className="flex gap-2 items-start"><span className="dot"/> <span>{x}</span></li>
))}
</ul>
<div className="mt-3 flex gap-2">
<button className="btn btn-primary" id="signup2">Создать аккаунт</button>
<button className="btn" id="signin2">Войти</button>
</div>
</div>


<div className="card p-6">
<h3 id="currency" className="mb-2 text-xl font-semibold">Виртуальная валюта</h3>
<p className="muted">Внутренняя валюта <b>C</b> живёт только на сайте. Реальные пополнения отключены по дизайну.</p>
<ul className="grid gap-2 mt-2">
{[
'Единица: Coin (C).',
'Источники: пятничные начисления и награды.',
'Прозрачная история транзакций.'
].map((x,i)=> (
<li key={i} className="flex gap-2 items-start"><span className="dot"/> <span>{x}</span></li>
))}
</ul>
</div>
</div>
</section>
);
}


export function FridayActions(){
return (
<section aria-labelledby="friday" className="py-6">
<div className="max-w-[1100px] mx-auto px-6 grid md:grid-cols-2 gap-4">
<div className="card p-6">
<h3 id="friday" className="mb-2 text-xl font-semibold">Пополнение по пятницам и награды</h3>
<p className="muted">Каждую пятницу активным пользователям приходит базовый объём монет + бонусы за полезные действия.</p>
<ul className="grid gap-2 mt-2">
{[
'Еженедельный дроп (фикс/диапазон).',
'Награды: вход, обучение, отзыв, победы.',
'Анти‑абуз: лимиты, анти‑бот, cooldown.'
].map((x,i)=> (
<li key={i} className="flex gap-2 items-start"><span className="dot"/> <span>{x}</span></li>
))}
</ul>
</div>
<div className="card p-6">
<h3 id="actions" className="mb-2 text-xl font-semibold">Примеры действий</h3>
<ul className="grid gap-2 mt-2">
{[
'Ежедневный вход — +3250 C (1 раз в сутки).',
'Пятничная получка — +22000 C (еженедельно).',
'Регистрация и туториал — +30000 C (однократно).',
'Отзыв/баг‑репорт — +15000 C (с модерацией).',
'Победа в мини‑игре — динамически.'
].map((x,i)=> (
<li key={i} className="flex gap-2 items-start"><span className="dot"/> <span>{x}</span></li>
))}
</ul>
</div>
</div>
</section>
);
}