/* eslint-disable @typescript-eslint/no-explicit-any */

import { Typography, Table } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../types/state';
import { ZoomInOutlined } from '@ant-design/icons';
import { getAllInvoices } from '../../../api/invoices/invoice/getAll';
import { Invoice } from '../../../types/invoice';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';
import InvoiceStatus from '../../status/InvoiceStatus';

const { Text, Link } = Typography;

const invoicesColumns = [
    {
        title: 'Company',
        dataIndex: 'company',
        key: 'company'
    },
    {
        title: 'Invoice name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'Discount',
        dataIndex: 'discount_percentage',
        key: 'discount_percentage'
    },
    {
        title: 'Price',
        dataIndex: 'total_price',
        key: 'total_price'
    },
    {
      title: 'Tax',
      dataIndex: 'total_tax',
      key: 'total_tax'
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status'
    },
    {
        title: 'Due date',
        dataIndex: 'due_date',
        key: 'due_date'
    },
    {
      title: '',
      dataIndex: 'operations',
      key: 'operations'
    },
  ];

const Invoices = () => {
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const [invoices, setInvoices] = useState<Array<Invoice>>([]);

    const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.company_name || 'Unknown';

    useEffect(() => {
        if (loggedInUserId) {
            getAllInvoices(loggedInUserId).then(response => {
                setInvoices(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId]);

      const invoiceData: Array<any> = useMemo(() => {
        const incomesListItem = invoices && invoices.map((invoice : Invoice) => {
        return {                    
            company: <Link href={`/external-company/${invoice.company_id}`}>{getVendorName(invoice.company_id)}</Link>,
            project: <Link href={`/project/${invoice.project_id}`}>{invoice.project_id}</Link>,
            subProject: <Link href={`/project/${invoice.sub_project_id}`}>{invoice.sub_project_id}</Link>,
            name: <Text>{invoice.invoice_display_name}</Text>,
            discount_percentage: <Text>{invoice.discount_percentage}%</Text>,
            total_price: <Text>{formatNumberWithSpaces(invoice.actual_price)} SEK</Text>,
            total_tax: <Text>{formatNumberWithSpaces(invoice.actual_tax)} SEK</Text>,
            status: <InvoiceStatus status={invoice.status} />,
            due_date: <Text>{invoice.due_date}</Text>,
            operations: <Link href={`/invoice/${invoice.id}`}><ZoomInOutlined /></Link>
          }
        })
        return incomesListItem;
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [invoices]);

    if (!invoiceData) return null;

    return  <Table size="small" columns={invoicesColumns} dataSource={invoiceData} />

}

export default Invoices;