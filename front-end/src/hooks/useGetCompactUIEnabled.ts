import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetCompactUIEnabled = () => {
    const compactUiThemeEnabled = useSelector((state : State) => state.user.settings.compact_ui);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => compactUiThemeEnabled, [compactUiThemeEnabled]);
}