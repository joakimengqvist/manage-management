import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetSubProjects = () => {
    const subProjects = useSelector((state : State) => state.application.subProjects);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => subProjects, [subProjects]);
}