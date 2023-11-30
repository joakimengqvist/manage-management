import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetExternalCompanies = () => {
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => externalCompanies, []);
}