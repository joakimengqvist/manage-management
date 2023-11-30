import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetLoggedInUserSettings = () => {
    const settings = useSelector((state : State) => state.user.settings);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => settings, [settings]);
}