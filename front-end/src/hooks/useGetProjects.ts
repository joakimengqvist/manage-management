import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetProjects = () => {
    const projects = useSelector((state : State) => state.application.projects);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => projects, []);
}