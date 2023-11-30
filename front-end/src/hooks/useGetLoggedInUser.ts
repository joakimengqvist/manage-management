import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetLoggedInUser = () => {
    const user = useSelector((state : State) => state.user);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => user, []);
}