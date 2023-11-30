/* eslint-disable @typescript-eslint/no-explicit-any */

import { Typography, Table } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { ZoomInOutlined } from '@ant-design/icons';
import { InvoiceItem } from '../../../interfaces/invoice';
import { getAllInvoiceItems } from '../../../api/invoices/invoiceItem/getAll';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';
import { useGetLoggedInUserId, useGetProducts } from '../../../hooks';

const { Text, Link } = Typography;

const invoiceItemsColumns = [
    {
        title: 'Product',
        dataIndex: 'product',
        key: 'product'
    },
    {
        title: 'Quantity',
        dataIndex: 'quantity',
        key: 'quantity'
    },
    {
        title: 'Total price',
        dataIndex: 'total_price',
        key: 'total_price'
    },
    {
        title: 'Tax',
        dataIndex: 'tax',
        key: 'tax'
    },
    {
        title: 'Discount',
        dataIndex: 'discount',
        key: 'discount'
    },
    {
      title: '',
      dataIndex: 'operations',
      key: 'operations'
    },
  ];

const InvoiceItems = () => {
    const loggedInUserId = useGetLoggedInUserId();
    const products = useGetProducts();
    const [invoiceItems, setInvoiceItems] = useState<Array<InvoiceItem>>([]);

    const getProductName = (id : string) => products?.[id]?.name;

    useEffect(() => {
        if (loggedInUserId) {
            getAllInvoiceItems(loggedInUserId).then(response => {
                setInvoiceItems(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId]);

      const invoiceItemsData: Array<any> = useMemo(() => {
        const invoiceIteamsListItem = invoiceItems && invoiceItems.map((invoiceItem : InvoiceItem) => {
        return {                    
            product: <Link href={`/product/${invoiceItem.product_id}`}>{getProductName(invoiceItem.product_id)}</Link>,
            quantity: <Text>{invoiceItem.quantity}</Text>,
            total_price: <Text>{formatNumberWithSpaces(invoiceItem.actual_price)}  SEK</Text>,
            tax: <Text>{formatNumberWithSpaces(invoiceItem.actual_tax)} SEK ({invoiceItem.tax_percentage}%)</Text>, 
            discount: <Text>{formatNumberWithSpaces(invoiceItem.original_price - invoiceItem.actual_price)} SEK ({invoiceItem.discount_percentage}%)</Text>,
            operations: <Link href={`/invoice-item/${invoiceItem.id}`}><ZoomInOutlined /></Link>
          }
        })
        return invoiceIteamsListItem;
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [invoiceItems]);

    if (!invoiceItemsData) return null;

    return  <Table size="small" columns={invoiceItemsColumns} dataSource={invoiceItemsData} />

}

export default InvoiceItems;