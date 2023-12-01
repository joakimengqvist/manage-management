/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Row, Table, Typography } from 'antd';
import { Card, notification } from 'antd';
import { getInvoiceItemById } from '../../../api/invoices/invoiceItem/getById';
import { useParams } from 'react-router-dom';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';
import { useGetLoggedInUser, useGetProducts, useGetUsers } from '../../../hooks';
import CreateNote from '../../notes/CreateNote';
import Notes from '../../notes/Notes';
import { NOTE_TYPE } from '../../../enums';
import { InvoiceItem, InvoiceItemNote } from '../../../interfaces';
import { getAllInvoiceItemNotesByInvoiceItemId } from '../../../api/notes/invoiceItem/getAllByInvoiceItemId';
import { createInvoiceItemNote } from '../../../api/notes/invoiceItem/create';

const { Title, Text, Link } = Typography;

const invoiceItemColumns : any = [
    {
        title: '',
        dataIndex: 'row_title',
        key: 'row_title',
        rowScope: 'row',
    },
    {
        title: 'Price',
        dataIndex: 'price',
        key: 'price',
    },
    {
        title: 'Tax',
        dataIndex: 'tax',
        key: 'tax',
    },
    {
      title: 'Discount',
      dataIndex: 'discount',
      key: 'discount',
    },
  ];


const InvoiceItemDetails = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const products = useGetProducts();
    const users = useGetUsers();
    const [invoiceItemNotes, setInvoiceItemNotes] = useState<Array<InvoiceItemNote> | null>(null);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [invoiceItem, setInvoiceItem] = useState<InvoiceItem>();
    const { id } =  useParams(); 
    const invoiceItemId = id || '';

    const getUserName = (id : string) => `${users?.[id]?.first_name} ${users?.[id]?.last_name}`;
    const getProductName = (id : string) => products?.[id]?.name;

    useEffect(() => {
        if (loggedInUser.id) {
            getInvoiceItemById(loggedInUser.id, invoiceItemId).then(response => {
                if (response?.error || !response?.data) {
                    api.error({
                        message: `Create project invoiceItem failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                        });
                    return
                }
                setInvoiceItem(response.data);
            })
            .catch(error => {
                api.error({
                    message: `Error creating project invoiceItem`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
            getAllInvoiceItemNotesByInvoiceItemId(loggedInUser.id, invoiceItemId).then(response => {
                setInvoiceItemNotes(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
    const onHandleNoteChange = (event : any) => setNote(event.target.value);

    const clearNoteFields = () => {
        setNoteTitle('');
        setNote('');
    }

    const onSubmitInvoiceItemNote = () => {
    const user = {
        id: loggedInUser.id,
        name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
        email: loggedInUser.email

    }
    createInvoiceItemNote(user, invoiceItemId, noteTitle, note).then((response) => {
        api.info({
            message: response.message,
            placement: "bottom",
            duration: 1.2,
        });
        }).catch(error => {
            api.error({
                message: `Error creating note`,
                description: error.toString(),
                placement: "bottom",
                duration: 1.4,
            });
        })
    }

    if (!invoiceItem) return null;

    const invoiceItemTableData = [{
        row_title:  <Text strong>Original (without discount)</Text>,
        price: <>{formatNumberWithSpaces(invoiceItem.original_price)} SEK</>,
        tax: <>{formatNumberWithSpaces(invoiceItem.original_tax)} SEK ({invoiceItem.tax_percentage}%)</>,
        discount: '-',
    },
    {
        row_title:  <Text strong>Actual (with discount)</Text>,
        price: <>{formatNumberWithSpaces(invoiceItem.actual_price)} SEK</>,
        tax: <>{formatNumberWithSpaces(invoiceItem.actual_tax)} SEK ({invoiceItem.tax_percentage}%)</>,
        discount: <>{formatNumberWithSpaces(invoiceItem.discount_amount)} SEK ({invoiceItem.discount_percentage}%)</>,
    }]

    return (
        <Row>
            {contextHolder}
            <Col span={16} style={{paddingRight: '8px'}}>
                <Card bodyStyle={{paddingTop: '12px'}}>
                    <Row>
                        <Col span={24}>
                        <Title level={5}>Invoice item</Title>
                        </Col>
                    </Row>
                    <Row>
                        <Col span={14} style={{padding: '12px 12px 12px 4px'}}>
                            <Text strong>Product</Text><br />
                            <Link href={`/product/${invoiceItem.product_id}`}>{getProductName(invoiceItem.product_id)}</Link><br />
                            <Text strong>Quantity</Text><br />
                            {invoiceItem.quantity}
                        </Col>
                        <Col span={5} style={{padding: '12px 12px 12px 0px'}}>
                            <Text strong>Created by</Text><br />
                            <Link href={`/user/${invoiceItem.created_by}`}>{getUserName(invoiceItem.created_by)}</Link><br />
                            <Text strong>Created at</Text><br />
                            {formatDateTimeToYYYYMMDDHHMM(invoiceItem.created_at)}<br />
                        </Col>
                        <Col span={5} style={{padding: '12px 12px 12px 0px'}}>
                            <Text strong>Updated by</Text><br />
                            <Link href={`/user/${invoiceItem.created_by}`}>{getUserName(invoiceItem.updated_by)}</Link><br />
                            <Text strong>Updated at</Text><br />
                            {formatDateTimeToYYYYMMDDHHMM(invoiceItem.updated_at)}
                        </Col>
                    </Row>
                    <Row>
                        <Col span={24} style={{marginTop: 12}}>
                        <Table
                            columns={invoiceItemColumns}
                            dataSource={invoiceItemTableData}
                            pagination={false}
                        />
                        </Col>
                    </Row>
                </Card>
            </Col>
            <Col span={8}>
                <Card>
                    <CreateNote
                        type={NOTE_TYPE.invoice_item}
                        title={noteTitle}
                        onTitleChange={onHandleNoteTitleChange}
                        note={note}
                        onNoteChange={onHandleNoteChange}
                        onClearNoteFields={clearNoteFields}
                        onSubmit={onSubmitInvoiceItemNote}
                    />
                    <Title level={4}>Notes</Title>
                    {invoiceItemNotes && invoiceItemNotes.length > 0 && 
                        <Notes notes={invoiceItemNotes} type={NOTE_TYPE.invoice_item} userId={loggedInUser.id} />
                    }
                </Card>
            </Col>
        </Row>
    );
};

export default InvoiceItemDetails;