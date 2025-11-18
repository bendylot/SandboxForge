import type { Metadata } from "next";
import "./colors.css";
import "./global.css";

export const metadata: Metadata = {
  title: "Casino — виртуальная валюта и мини‑игры",
  description:
    "Регистрация, виртуальная валюта, пятничные начисления, мини‑игры.",
  icons: { icon: "/favicon.ico" },
  openGraph: {
    title: "Casino — виртуальная валюта и мини‑игры",
    description: "Без реальных денег. Еженедельные начисления и награды.",
    type: "website",
    images: ["/og.png"],
  },
  themeColor: "#0ea5e9",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ru">
      <body className="bg-[#0b1020] text-white leading-relaxed">
        {children}
      </body>
    </html>
  );
}
