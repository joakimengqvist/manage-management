import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetInvoices = () => {
    const invoices = useSelector((state : State) => state.application.invoices);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => invoices, [invoices]);
}