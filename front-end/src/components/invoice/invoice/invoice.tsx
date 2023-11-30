/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Typography, Row, Col, Button, Table, notification } from 'antd';
import { useEffect, useState } from 'react';
import { getInvoiceById } from '../../../api/invoices/invoice/getById';
import { Invoice, InvoiceItem } from '../../../interfaces/invoice'
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';
import { GoldTag } from '../../tags/GoldTag';
import { GreenTag } from '../../tags/GreenTag';
import InvoiceStatus from '../../status/InvoiceStatus';
import { useGetExternalCompanies, useGetLoggedInUserId, useGetProducts, useGetProjects, useGetUsers } from '../../../hooks';
import { useGetSubProjects } from '../../../hooks/useGetSubProjects';
import { getInvoiceItemsByIds } from '../../../api/invoices/invoiceItem/getAllByIds';
// import UpdateInvoice from './updateInvoice';

const { Text, Title, Link } = Typography;

const invoiceItemsColumns = [
    {
        title: 'Product',
        dataIndex: 'product_id',
        key: 'product_id',
    },
    {
        title: 'Quantity',
        dataIndex: 'quantity',
        key: 'quantity',
    },
    {
        title: 'Discount',
        dataIndex: 'discount_percentage',
        key: 'discount_percentage',
    },
    {
        title: 'Tax',
        dataIndex: 'actual_tax',
        key: 'actual_tax',
    },
    {
        title: 'Price',
        dataIndex: 'actual_price',
        key: 'actual_price',
    },
]

const InvoiceDetails = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const users = useGetUsers();
    const projects = useGetProjects();
    const subProjects = useGetSubProjects();
    const products = useGetProducts();
    const externalCompanies = useGetExternalCompanies();
    const [invoice, setInvoice] = useState<null | Invoice>(null);
    const [editing, setEditing] = useState(false);
    const [invoiceItemsTableData, setInvoiceItemsTableData] = useState<any>([]);
    const { id } =  useParams(); 
    const invoiceId = id || '';

    useEffect(() => {
        if (loggedInUserId) {
            getInvoiceById(loggedInUserId, invoiceId).then(response => {
                setInvoice(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
    }, [loggedInUserId]);

    useEffect(() => {
        if (invoice) {
            getInvoiceItemsByIds(loggedInUserId, invoice.invoice_items).then(response => {
                if (response?.error) {
                    api.error({
                        message: 'Error',
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                    });
                    return;
                }

                const tableData = response.data.map((invoiceItem : InvoiceItem) => {
                    return {
                        product_id: getProductName(invoiceItem?.product_id || ''),
                        quantity: invoiceItem?.quantity,
                        discount_percentage: `${invoiceItem?.discount_percentage}%`,
                        actual_tax: `${invoiceItem?.actual_tax} SEK`,
                        actual_price: `${invoiceItem?.actual_price} SEK`,
                    }});

                setInvoiceItemsTableData(tableData);
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
    }, [invoice])

    const getUserName = (id : string) => `${users?.[id]?.first_name} ${users?.[id]?.last_name}`;
    const getVendorName = (id : string) => externalCompanies?.[id]?.company_name;
    const getProjectName = (id : string) => projects?.[id]?.name;
    const getProductName = (id : string) => products?.[id]?.name;

    const getSubProjectName = (id : string) => subProjects.find(subProject => subProject.id === id)?.name;

    return (<>
        {contextHolder}
        {invoice && (
            <Row>
                <Col span={16} style={{paddingRight: '24px'}}>
                        <Card style={{marginBottom: '16px'}}>
                        {editing ? (<></>
                        // <UpdateInvoice invoice={invoice} setEditing={setEditing} />
                        ) : (
                        <>
                            <Row>
                                <Col span={24} style={{marginBottom: '0px'}}>
                                    <div style={{display: 'flex', justifyContent: 'space-between'}}>
                                        <div style={{paddingBottom: '16px'}}>
                                            <Title level={4} style={{marginBottom: '0px'}}>{invoice.invoice_display_name}</Title>
                                            <Link style={{fontSize: '16px'}}href={`/external-company/${invoice.company_id}`}>{getVendorName(invoice.company_id)}</Link><br />
                                            <Link style={{fontSize: '16px'}}href={`/project/${invoice.project_id}`}>{getProjectName(invoice.project_id)}</Link><br />
                                            {invoice.invoice_description}<br />
                                            {formatDateTimeToYYYYMMDDHHMM(invoice.invoice_date)}<br />
                                        </div>
                                            <div style={{display: 'flex', flexDirection: 'column', alignItems: 'end', paddingRight: '20px', gap: '4px'}}>
                                                <Button onClick={() => setEditing(true)} style={{marginTop: '16px'}}>Edit invoice info</Button>
                                                <div style={{display: 'flex', justifyContent: 'flex-end', gap: '4px'}}>
                                                    {invoice.paid && <GreenTag label="Paid" />}
                                                    {invoice.statistics_invoice && <GoldTag label="Statistics invoice" />}
                                                </div>
                                                <Link href={`/income/${invoice.income_id}`}>Go to income</Link><br/>
                                            </div>
                                    </div>
                                </Col>
                            </Row>
                            <Card style={{marginRight: '20px', paddingLeft: '20px'}}>
                                <Row>       
                                    <Col span={8}  style={{padding: '0px 12px 0px 0px'}}>
                                        {invoice.project_id && (<>
                                            <Text strong>Project</Text><br/>
                                            <Link href={`/project/${invoice.project_id}`}>{getProjectName(invoice.project_id)}</Link><br/>
                                        </>)}
                                        {invoice.sub_project_id && (<>
                                            <Text strong>Sub Project</Text><br/>
                                            <Link href={`/sub-project/${invoice.sub_project_id}`}>{getSubProjectName(invoice.sub_project_id)}</Link><br/>
                                        </>)}
                                        <Text strong>Invoice status</Text><br/>
                                        <InvoiceStatus status={invoice.status} /><br/><br/>
                                    </Col>
                                    <Col span={8}  style={{padding: '0px 12px 0px 0px'}}>
                                        <Text strong>Invoice date</Text><br/>
                                        {formatDateTimeToYYYYMMDDHHMM(invoice.invoice_date)}<br/>
                                        <Text strong>Due date</Text><br/>
                                        {formatDateTimeToYYYYMMDDHHMM(invoice.due_date)}<br/>
                                    </Col>
                                    <Col span={8} style={{padding: '0px 12px 0px 0px'}}>
                                        <Text strong>Created by</Text><br/>
                                        <Link href={`/user/${invoice.created_by}`}>{getUserName(invoice.created_by)}</Link><br/>
                                        <Text strong>Created at</Text><br/>
                                        {formatDateTimeToYYYYMMDDHHMM(invoice.created_at)}<br/>
                                        <Text strong>Modified by</Text><br/>
                                        <Link href={`/user/${invoice.updated_by}`}>{getUserName(invoice.updated_by)}</Link><br/>
                                        <Text strong>Modified at</Text><br/>
                                        {formatDateTimeToYYYYMMDDHHMM(invoice.updated_at)}<br/>
                                    </Col>
                                </Row>
                            </Card>
                            <Row>
                            <Col span={24} style={{paddingRight: '20px', marginTop: '20px'}}>
                                <Table size="small" dataSource={invoiceItemsTableData} columns={invoiceItemsColumns} pagination={false} />
                                    <div style={{display: 'flex', justifyContent: 'flex-end', alignContent: 'center', gap: '24px', marginTop: '28px', paddingRight: '40px'}}>
                                        <div>
                                            <Text strong>Discount percentage</Text><br/>
                                            {invoice.discount_percentage}%<br/>
                                        </div>
                                        <div>
                                            <Text strong>Total amount</Text><br/>
                                            {formatNumberWithSpaces(invoice.actual_price)} SEK
                                        </div>
                                        <div>
                                            <Text strong>Total tax</Text><br/>
                                            {formatNumberWithSpaces(invoice.actual_tax)} SEK
                                        </div>
                                    </div>
                                </Col>
                            </Row>
                        </>)}
                        </Card>
                </Col>
                <Col span={8}>
                </Col>
            </Row>
        )}
    </>)
}

export default InvoiceDetails;