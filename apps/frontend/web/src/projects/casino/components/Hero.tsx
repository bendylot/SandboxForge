export default function Hero(){
    return (
        <section className="py-14">
            <div className="max-w-[1100px] mx-auto px-6 grid md:grid-cols-2 gap-6">
                <div className="card p-6">
                    <h1 className="text-[clamp(28px,4vw,44px)] leading-tight font-bold mb-3">
                        Виртуальная валюта. Мини-игры.<br/>
                        <span className="bg-clip-text text-transparent"
                            style={{backgroundImage:'linear-gradient(135deg, #0ea5e9, #22d3ee)'}}>
                            Без реальных денег
                        </span>.
                    </h1>
                    <p className="muted text-lg mb-4">Регистрируйся, получай еженедельное пополнение по пятницам и трать монеты в мини-играх. Плюс призы за активность.</p>
                    <div className="flex gap-2 flex-wrap mb-3">
                        {['Регистрация','Еженедельное пополнение','Награды за действия','Единый пул мини-игр'].map(b=>
                        <span key={b} className="badge">{b}</span>
                        )}
                    </div>
                    <div className="flex gap-2 mt-2">
                        <button className="btn btn-primary" id="cta-play">Начать</button>
                        <a className="btn" href="#games">Посмотреть мини-игры</a>
                    </div>
                </div>


                <aside className="card p-6">
                <h2 className="mb-2 font-semibold text-xl">Как это работает</h2>
                <ul className="grid gap-2">
                {[
                <><b>1. Регистрация.</b> Приветственный пакет виртуальной валюты.</>,
                <><b>2. Пополнение по пятницам.</b> Еженедельно + бонусы.</>,
                <><b>3. Единый баланс.</b> Один кошелёк — множество мини‑игр.</>,
                <><b>4. Без реальных денег.</b> Только фан и геймификация.</>
                ].map((x,i)=> (
                <li key={i} className="flex gap-2 items-start"><span className="dot"/> <span>{x}</span></li>
                ))}
                </ul>
                <p className="mt-3 border-l-4 border-warn/100 rounded-md p-3 text-[#ffe9c2]"
                style={{background:'rgba(245, 158, 11, .08)'}}>Это не азартные игры на деньги. Валюта не продаётся и не обменивается.</p>
                </aside>
            </div>
        </section>
    );
}