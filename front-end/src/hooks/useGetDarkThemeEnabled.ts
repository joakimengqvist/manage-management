import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetDarkThemeEnabled = () => {
    const darkThemeEnabled = useSelector((state : State) => state.user.settings.dark_theme);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => darkThemeEnabled, [darkThemeEnabled]);
}