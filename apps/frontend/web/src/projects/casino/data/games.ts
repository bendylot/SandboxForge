export type Game = {
    id: string;
    title: string;
    kind: "Аркада" | "Память" | "Рефлекс";
    price: string;
    description: string;
};


export const games: Game[] = [
    { id: "dice", title: "Dice Rush", kind: "Аркада", price: "2 C", description: "Броски кубика, риск/награда." },
    { id: "cards", title: "Card Match", kind: "Память", price: "3 C", description: "Открой пары за минимум ходов." },
    { id: "tower", title: "Tower Stack", kind: "Рефлекс", price: "5 C", description: "Строй башню точными таймингами." }
];