import type { Config } from "tailwindcss";


export default {
    content: [
        "./app/**/*.{ts,tsx}",
        "./components/**/*.{ts,tsx}"
    ],
    theme: {
        extend: {
            colors: {
                bg: "#0b1020",
                card: "#121936",
                muted: "#97a0c3",
                brand: "#0ea5e9",
                brand2: "#22d3ee",
                ok: "#22c55e",
                bad: "#ef4444",
                warn: "#f59e0b"
            },
            boxShadow: {
                soft: "0 10px 30px rgba(2,6,23,.45)"
            },
            borderRadius: {
                soft: "16px"
            }
        }
    },
    plugins: []
} satisfies Config;