import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetProducts = () => {
    const products = useSelector((state : State) => state.application.products);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => products, []);
}