import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetPrivileges = () => {
    const privileges = useSelector((state : State) => state.application.privileges);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => privileges, [privileges]);
}