import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetLoggedInUserId = () => {
    const userId = useSelector((state : State) => state.user.id);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => userId, [userId]);
}