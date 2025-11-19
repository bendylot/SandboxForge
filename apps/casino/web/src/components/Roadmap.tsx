export default function Roadmap(){
return (
<section aria-labelledby="roadmap" className="py-6">
<div className="max-w-[1100px] mx-auto px-6 grid md:grid-cols-3 gap-4">
{[
{title:'Дорожная карта', items:['v0: Лендинг, регистрация, кошелёк, пятничные начисления.','v1: 3 мини‑игры, лидеры, ачивки.','v2: Сезоны, задачи недели, кланы.']},
{title:'Прозрачность', items:['Публичная история транзакций и правила наград.','Ограничения на фарм и анти‑бот политика.']},
{title:'Безопасность', items:['Хранение паролей с хэшированием.','Лимиты запросов и защита от накруток.']},
].map((col,i)=> (
<div key={i} className="card p-6">
<h3 id={i===0? 'roadmap': undefined} className="mb-2 text-xl font-semibold">{col.title}</h3>
<ul className="grid gap-2">
{col.items.map((x,idx)=> (
<li key={idx} className="flex gap-2 items-start"><span className="dot"/> <span>{x}</span></li>
))}
</ul>
</div>
))}
</div>
</section>
);
}