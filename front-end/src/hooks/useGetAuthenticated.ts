import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetAuthenticated = () => {
    const authenticated = useSelector((state : State) => state.user.authenticated);
    return useMemo(() => authenticated, [authenticated]);
}