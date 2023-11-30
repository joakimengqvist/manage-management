import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetLoggedInUserPrivileges = () => {
    const privileges = useSelector((state : State) => state.user.privileges);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => privileges, [privileges]);
}