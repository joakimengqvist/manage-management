import { useMemo } from "react";
import { useSelector } from "react-redux";
import { State } from "../interfaces";

export const useGetInvoiceItems = () => {
    const invoiceItems = useSelector((state : State) => state.application.invoiceItems);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return useMemo(() => invoiceItems, [invoiceItems]);
}