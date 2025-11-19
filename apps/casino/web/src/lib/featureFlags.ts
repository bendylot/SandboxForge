export type Flags = {
    showGames: boolean;
    showRoadmap: boolean;
    allowOAuth: boolean;
};


export function getFlags(): Flags {
    return {
        showGames: process.env.NEXT_PUBLIC_SHOW_GAMES !== "false",
        showRoadmap: process.env.NEXT_PUBLIC_SHOW_ROADMAP !== "false",
        allowOAuth: process.env.NEXT_PUBLIC_ALLOW_OAUTH !== "false",
    };
}