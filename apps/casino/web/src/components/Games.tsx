import { games } from "../data/games";

export default function Games() {
  return (
    <section id="games" aria-labelledby="pool" className="py-6">
      <div className="max-w-[1100px] mx-auto px-6">
        <div className="card p-6">
          <h2 id="pool" className="text-xl font-semibold mb-2">
            Единый пул мини‑игр
          </h2>
          <p className="muted mb-3">
            Игры используют общий кошелёк. Для запуска/хода тратятся монеты.
            Баланс и транзакции едины.
          </p>

          <div className="grid md:grid-cols-3 gap-4" role="list">
            {games.map((g) => (
              <article
                key={g.id}
                className="card p-4 grid gap-2"
                role="listitem"
              >
                <div className="flex items-center justify-between">
                  <h3 className="m-0 font-semibold">{g.title}</h3>
                  <span className="label">{g.kind}</span>
                </div>
                <p className="muted">{g.description}</p>
                <div className="muted">
                  Цена: <b>{g.price}</b>
                </div>
                <div className="flex gap-2">
                  <button className="btn btn-sm" data-game={g.id}>
                    Правила
                  </button>
                  <button className="btn btn-sm btn-primary" data-play={g.id}>
                    Играть
                  </button>
                </div>
              </article>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}
