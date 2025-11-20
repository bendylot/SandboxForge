"use client";
import { useEffect, useRef, useState } from "react";

export default function ModalSignup() {
  const [open, setOpen] = useState(false);
  const dialogRef = useRef<HTMLDivElement>(null);

  // Привязываем внешние кнопки по id (как в исходнике)
  useEffect(() => {
    const a = document.getElementById("btn-signup");
    const b = document.getElementById("signup2");
    const on = () => setOpen(true);
    a?.addEventListener("click", on);
    b?.addEventListener("click", on);
    return () => {
      a?.removeEventListener("click", on);
      b?.removeEventListener("click", on);
    };
  }, []);

  useEffect(() => {
    function onKey(e: KeyboardEvent) {
      if (e.key === "Escape") setOpen(false);
    }
    if (open) document.addEventListener("keydown", onKey);
    return () => document.removeEventListener("keydown", onKey);
  }, [open]);

  const err = (name: string, msg = "") => {
    const el = dialogRef.current?.querySelector(
      `[data-err-for="${name}"]`
    ) as HTMLElement | null;
    if (el) el.textContent = msg;
  };

  async function onSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const form = e.currentTarget as HTMLFormElement & {
      nick: HTMLInputElement;
      email: HTMLInputElement;
      password: HTMLInputElement;
    };
    const nick = form.nick.value.trim();
    const email = form.email.value.trim();
    const password = form.password.value;
    const terms = (document.getElementById("signup-terms") as HTMLInputElement)
      ?.checked;

    ["nick", "email", "password", "terms"].forEach((n) => err(n, ""));
    let ok = true;
    const emailRe = /^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/i;
    if (nick.length < 4) {
      err("nick", "Минимум 4 символа");
      ok = false;
    }
    if (!emailRe.test(email)) {
      err("email", "Неверный email");
      ok = false;
    }
    if (!password || password.length < 4) {
      err("password", "Минимум 4 символов");
      ok = false;
    }
    if (!terms) {
      err("terms", "Нужно принять правила");
      ok = false;
    }
    if (!ok) return;

    try {
      const res = await fetch("/api/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ nick, email, password }),
      });
      if (!res.ok) {
        const data = await res.json().catch(() => ({ message: "Ошибка" }));
        throw new Error(data.message || "Ошибка регистрации");
      }
      setOpen(false);
      alert("Готово! Проверь почту для подтверждения.");
    } catch (e: any) {
      err("email", e.message || "Не удалось зарегистрироваться");
    }
  }

  if (!open) return null;
  return (
    <div
      className="fixed inset-0 z-[1000]"
      aria-hidden={false}
      role="dialog"
      aria-modal="true"
    >
      <div
        className="absolute inset-0 bg-[#030712]/85 backdrop-blur-[6px]"
        onClick={() => setOpen(false)}
      />
      <div
        ref={dialogRef}
        tabIndex={-1}
        role="document"
        className="relative mx-auto mt-[6vh] max-w-[560px] p-0 rounded-2xl border border-white/15 bg-[#0a1428]/95 text-white shadow-[0_25px_80px_rgba(2,6,23,0.8)]"
      >
        <header className="px-5 py-4 border-b border-white/10 flex items-center justify-between bg-white/5 rounded-t-2xl">
          <h2 className="text-lg font-semibold">Регистрация</h2>
          <button
            className="text-muted text-2xl"
            aria-label="Закрыть"
            onClick={() => setOpen(false)}
          >
            ×
          </button>
        </header>
        <section className="px-5 pt-4 pb-2 space-y-4">
          <p className="text-white/90">
            Выберите способ регистрации. Реальных денег нет — только виртуальные
            монеты и фан.
          </p>
          <div className="grid md:grid-cols-3 gap-2">
            {[
              { k: "google", icon: "🟢", t: "Google" },
              { k: "yandex", icon: "🟥", t: "Яндекс" },
              { k: "vk", icon: "🟦", t: "ВК" },
            ].map((x) => (
              <button
                key={x.k}
                className="btn btn-sm w-full"
                onClick={() => (location.href = `/auth/${x.k}`)}
                aria-label={`Регистрация через ${x.t}`}
              >
                <span className="opacity-90 text-sm">{x.icon}</span> {x.t}
              </button>
            ))}
          </div>
          <div className="grid place-items-center">
            <span className="badge text-[12px] bg-white/10 text-white">
              или
            </span>
          </div>
          <form
            onSubmit={onSubmit}
            className="grid gap-3 bg-white/5 rounded-xl p-4 border border-white/10"
            noValidate
          >
            <div className="grid gap-1">
              <label htmlFor="signup-nick" className="font-semibold text-white">
                Никнейм
              </label>
              <input
                id="signup-nick"
                name="nick"
                autoComplete="nickname"
                required
                minLength={4}
                maxLength={15}
                placeholder="Например, bendylot"
                className="bg-white/10 border border-white/20 rounded-lg px-3 py-2 outline-none focus:border-brand/60 placeholder:text-white/60"
              />
              <small
                className="text-[#ffb4b4] min-h-[1em]"
                data-err-for="nick"
              />
            </div>
            <div className="grid gap-1">
              <label
                htmlFor="signup-email"
                className="font-semibold text-white"
              >
                Email
              </label>
              <input
                id="signup-email"
                name="email"
                type="email"
                autoComplete="email"
                required
                placeholder="you@example.com"
                className="bg-white/10 border border-white/20 rounded-lg px-3 py-2 outline-none focus:border-brand/60 placeholder:text-white/60"
              />
              <small
                className="text-[#ffb4b4] min-h-[1em]"
                data-err-for="email"
              />
            </div>
            <div className="grid gap-1">
              <label htmlFor="signup-pass" className="font-semibold text-white">
                Пароль
              </label>
              <input
                id="signup-pass"
                name="password"
                type="password"
                autoComplete="new-password"
                required
                minLength={4}
                placeholder="Минимум 4 символов"
                className="bg-white/10 border border-white/20 rounded-lg px-3 py-2 outline-none focus:border-brand/60 placeholder:text-white/60"
              />
              <small
                className="text-[#ffb4b4] min-h-[1em]"
                data-err-for="password"
              />
            </div>
            <label className="inline-flex gap-2 items-center select-none text-white">
              <input type="checkbox" id="signup-terms" required />
              <span>Согласен с правилами и обработкой данных</span>
            </label>
            <small
              className="text-[#ffb4b4] min-h-[1em]"
              data-err-for="terms"
            />
            <div className="flex gap-2 mt-1">
              <button type="submit" className="btn btn-primary">
                Зарегистрироваться
              </button>
              <button
                type="button"
                className="btn"
                onClick={() => setOpen(false)}
              >
                Отмена
              </button>
            </div>
          </form>
        </section>
        <footer className="px-5 py-3 border-t border-white/10 text-white/70 bg-white/5 rounded-b-2xl">
          Виртуальная валюта не имеет денежной стоимости.
        </footer>
      </div>
    </div>
  );
}
