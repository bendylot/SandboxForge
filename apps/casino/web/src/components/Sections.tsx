export function RegCurrency() {
  return (
    <section aria-labelledby="reg" className="py-6">
      <div className="max-w-[1100px] mx-auto px-6 grid md:grid-cols-2 gap-4">
        <div className="card p-6">
          <h3 id="reg" className="mb-2 text-xl font-semibold">
            Виртуальная валюта
          </h3>
          <p className="muted">
            Внутренняя валюта <b>C</b> живёт только на сайте. Реальные
            пополнения отключены по дизайну.
          </p>
          <ul className="grid gap-2 mt-2">
            {[
              "Единица: Coin (C).",
              "Источники: пятничные начисления и награды.",
              "Прозрачная история транзакций.",
            ].map((x, i) => (
              <li key={i} className="flex gap-2 items-start">
                <span className="dot" /> <span>{x}</span>
              </li>
            ))}
          </ul>
        </div>
        <div className="card p-6">
          <h3 className="mb-2 font-semibold text-xl">Как это работает</h3>
          <ul className="grid gap-2">
            {[
              <>
                <b>1. Регистрация.</b> Приветственный пакет виртуальной валюты.
              </>,
              <>
                <b>2. Пополнение по пятницам.</b> Еженедельно + бонусы.
              </>,
              <>
                <b>3. Единый баланс.</b> Один кошелёк — множество мини‑игр.
              </>,
              <>
                <b>4. Без реальных денег.</b> Только фан и геймификация.
              </>,
            ].map((x, i) => (
              <li key={i} className="flex gap-2 items-start">
                <span className="dot" /> <span>{x}</span>
              </li>
            ))}
          </ul>
          <p
            className="mt-3 border-l-4 border-warn/100 rounded-md p-3 text-[#ffe9c2]"
            style={{ background: "rgba(245, 158, 11, .08)" }}
          >
            Это не азартные игры на деньги. Валюта не продаётся и не обменивается.
          </p>
          <div className="mt-4 flex flex-wrap gap-3 items-center">
            <button type="button" className="btn btn-primary" id="signup2">
              Зарегистрироваться
            </button>
            <span className="muted text-sm">
              Получи приветственный пакет и пятничные начисления.
            </span>
          </div>
        </div>
      </div>
    </section>
  );
}

export function FridayActions() {
  return (
    <section aria-labelledby="friday" className="py-6">
      <div className="max-w-[1100px] mx-auto px-6 grid md:grid-cols-2 gap-4">
        <div className="card p-6">
          <h3 id="friday" className="mb-2 text-xl font-semibold">
            Пополнение по пятницам и награды
          </h3>
          <p className="muted">
            Каждую пятницу активным пользователям приходит базовый объём монет +
            бонусы за полезные действия.
          </p>
          <ul className="grid gap-2 mt-2">
            {[
              "Еженедельный дроп (фикс/диапазон).",
              "Награды: вход, обучение, отзыв, победы.",
              "Анти‑абуз: лимиты, анти‑бот, cooldown.",
            ].map((x, i) => (
              <li key={i} className="flex gap-2 items-start">
                <span className="dot" /> <span>{x}</span>
              </li>
            ))}
          </ul>
        </div>
        <div className="card p-6">
          <h3 id="actions" className="mb-2 text-xl font-semibold">
            Примеры действий
          </h3>
          <ul className="grid gap-2 mt-2">
            {[
              "Ежедневный вход — +3250 C (1 раз в сутки).",
              "Пятничная получка — +22000 C (еженедельно).",
              "Регистрация и туториал — +30000 C (однократно).",
              "Отзыв/баг‑репорт — +15000 C (с модерацией).",
              "Победа в мини‑игре — динамически.",
            ].map((x, i) => (
              <li key={i} className="flex gap-2 items-start">
                <span className="dot" /> <span>{x}</span>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </section>
  );
}
