export default function Header(){
  return (
    <header className="sticky top-0 z-20 border-b border-white/10 bg-[#0b1020]/50 backdrop-blur">
      <div className="mx-auto flex max-w-[1100px] items-center justify-between gap-4 p-6">
        <a
          className="flex items-center gap-2 font-bold tracking-wide"
          href="#"
          aria-label="Casino — главная"
        >
          <span
            className="grid h-[34px] w-[34px] place-items-center rounded-[10px] shadow-soft text-[#061226] font-black"
            style={{ backgroundImage: "linear-gradient(135deg, #0ea5e9, #22d3ee)" }}
          >
            C
          </span>
          <span>Casino</span>
        </a>

        <nav className="flex gap-2" aria-label="Действия">
          <button type="button" className="btn btn-sm btn-ghost" id="btn-signin">
            Войти
          </button>
          <button type="button" className="btn btn-sm btn-primary" id="btn-signup">
            Зарегистрироваться
          </button>
        </nav>
      </div>
    </header>
  );
}
