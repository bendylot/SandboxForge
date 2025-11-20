export default function Hero() {
  return (
    <section className="py-14">
      <div className="max-w-[1100px] mx-auto px-6">
        <div className="space-y-4">
          <h1 className="text-[clamp(28px,4vw,44px)] leading-tight font-bold mb-3">
            Виртуальная валюта. Мини-игры.
            <br />
            <span
              className="bg-clip-text text-transparent"
              style={{
                backgroundImage: "linear-gradient(135deg, #0ea5e9, #22d3ee)",
              }}
            >
              Без реальных денег
            </span>
            .
          </h1>
          <p className="muted text-lg mb-4">
            Регистрируйся, получай еженедельное пополнение по пятницам и трать
            монеты в мини-играх. Плюс призы за активность.
          </p>
          <div className="flex gap-2 flex-wrap mb-3">
            {[
              "Регистрация",
              "Еженедельное пополнение",
              "Награды за действия",
              "Единый пул мини-игр",
            ].map((b) => (
              <span key={b} className="badge">
                {b}
              </span>
            ))}
          </div>
          <div className="flex gap-2 mt-2">
            <button className="btn btn-primary" id="cta-play">
              Начать
            </button>
            <a className="btn" href="#games">
              Посмотреть мини-игры
            </a>
          </div>
        </div>
      </div>
    </section>
  );
}
