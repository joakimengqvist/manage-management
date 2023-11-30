import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetUsers = () => {
    const users = useSelector((state : State) => state.application.users);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => users, []);
}