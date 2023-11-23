export type Invoice = {
    id: string
    company_id: string
    project_id: string
    sub_project_id: string
    income_id: string
    invoice_display_name: string
    invoice_description: string
    statistics_invoice: boolean
    invoice_items: Array<string>
    original_price: number
    actual_price: number
    discount_percentage: number
    discount_amount: number
    original_tax: number
    actual_tax: number
    invoice_date: string
    due_date: string
    paid: boolean
    status: string
    payment_date: string
    created_by: string
    created_at: string
    updated_by: string
    updated_at: string
}

export type InvoiceItem = {
    id: string
    product_id: string
    quantity: number
    original_price: number
    actual_price: number
    discount_percentage: number
    discount_amount: number
    tax_percentage: number
    original_tax: number
    actual_tax: number
    created_by: string
    created_at: string
    updated_by: string
    updated_at: string
}