export const GAMEMODE_RANKED = "Ranked" as const;
export const GAMEMODE_UNRANKED = "Unranked" as const;

export const GAMEMODES = [GAMEMODE_RANKED, GAMEMODE_UNRANKED] as const;

export const GAME_FORMAT_BO3 = "Bo3" as const;
export const GAME_FORMAT_BO5 = "Bo5" as const;
export const GAME_FORMAT_ANY = "Any" as const;

export const GAME_FORMATS = [GAME_FORMAT_BO3, GAME_FORMAT_BO5, GAME_FORMAT_ANY] as const;

export type GameMode = typeof GAMEMODES[number];
export type GameFormat = typeof GAME_FORMATS[number];

export const GAME_FORMATS_BY_GAME_MODE: Record<GameMode, readonly GameFormat[]> = {
    [GAMEMODE_RANKED]: [GAME_FORMAT_BO3],
    [GAMEMODE_UNRANKED]: [GAME_FORMAT_BO3, GAME_FORMAT_BO5, GAME_FORMAT_ANY],
} as const;
