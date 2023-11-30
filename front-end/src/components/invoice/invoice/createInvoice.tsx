/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Checkbox, Col, Row, Table, Typography } from 'antd';
import { Button, Input, Space, Card, notification, DatePicker, Select } from 'antd';
import { createInvoice } from '../../../api/invoices/invoice/create';
import { CloseCircleOutlined } from '@ant-design/icons';
import { InvoiceItem } from '../../../interfaces/invoice';
import { statusOptions } from './options';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';
import { getInvoiceItemsByIds } from '../../../api/invoices/invoiceItem/getAllByIds';
import { useGetExternalCompanies, useGetInvoiceItems, useGetLoggedInUserId, useGetProducts, useGetProjects } from '../../../hooks';
import { useGetSubProjects } from '../../../hooks/useGetSubProjects';

const { Text, Title } = Typography;
const { TextArea } = Input;

const invoiceItemsColumns = [
    {
        title: 'Product',
        dataIndex: 'product_name',
        key: 'product_name'
    },
    {
        title: 'Quantity',
        dataIndex: 'quantity',
        key: 'quantity'
    },
    {
        title: 'Discount',
        dataIndex: 'discount',
        key: 'discount'
    },
    {
      title: 'Price',
      dataIndex: 'actual_price',
      key: 'actual_price'
    },
    {
        title: 'tax',
        dataIndex: 'actual_tax',
        key: 'actual_tax'
    },
    {
      title: '',
      dataIndex: 'operations',
      key: 'operations'
    },
  ];

const CreateInvoice = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const projects = useGetProjects();
    const subProjects = useGetSubProjects();
    const products = useGetProducts();
    const externalCompanies = useGetExternalCompanies();
    const allInvoiceItems = useGetInvoiceItems();

    const [companyId, setCompanyId] = useState('');
    const [projectId, setProjectId] = useState('');
    const [subProjectId, setSubProjectId] = useState('');
    const [invoiceDisplayName, setInvoiceDisplayName] = useState('');
    const [invoiceDescription, setInvoiceDescription] = useState('');
    const [isStatisticsInvoice, setIsStatisticsInvoice] = useState(false);
    const [invoiceItems, setInvoiceItems] = useState([]);
    const [selectedInvoiceItemsDetails, setSelectedInvoiceItemsDetails]  = useState<any>([]);
    const [discountPercentage, setDiscountPercentage] = useState(0);
    const [invoiceDate, setInvoiceDate] = useState('');
    const [dueDate, setDueDate] = useState('');
    const [status, setStatus] = useState('');
    const [paymentDate, setPaymentDate] = useState('');
    const [subProjectOptions, setSubProjectOptions] = useState<any>([]);

    const getProductName = (id : string) => products?.[id]?.name;

    const onChangeInvoiceDate = (value : any) => {
        if (value) {
            setInvoiceDate(value.$d)
        }
    }

    const onChangeDueDate = (value : any) => {
        if (value) {
            setDueDate(value.$d)
        }
    }

    const onChangePaymentDate = (value : any) => {
        if (value) {
            setPaymentDate(value.$d)
        }
    }

    const onChangeCompany = (value : string) =>  setCompanyId(value); 
    const onChangeProject = (value: string) => {
        setProjectId(value)

        let subProjectsInProject: Array<string> = [];
        Object.keys(projects).forEach(projectId => {
            if (projects[projectId].id === value) {
                subProjectsInProject = projects[projectId].sub_projects;
            }
        })

        const subProjectOptionsArray : Array<any> = [];

        subProjects.forEach((subProject : any) => {
            if (subProjectsInProject.includes(subProject.id)) {
                subProjectOptionsArray.push({
                    label: subProject.name,
                    value: subProject.id
                })
            }
        })

        setSubProjectOptions(subProjectOptionsArray);
    }
    const onChangeSubProject = (value: string) => setSubProjectId(value);
    const onChangeStatus = (value: string) => setStatus(value);

    const onChangeInvoiceItems = (value : Array<never>) => {
        setInvoiceItems(value)
        getInvoiceItemsByIds(loggedInUserId, value).then((response : any) => {
            if (response?.error) {
                api.error({
                    message: `Create external company failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                });
                return;
            }

            const selectedItems = response.data.map((item : any) => {      
                return {
                    product_name: getProductName(item.product_id),
                    quantity: item.quantity,
                    discount: `${item.discount_percentage}%`,
                    actual_price: item.actual_price,
                    actual_tax: item.actual_tax,
                    operations: (
                        <Space direction="horizontal">
                            <CloseCircleOutlined style={{ color: 'red' }} onClick={() => onRemoveInvoiceItem(item.id)} />
                        </Space>)
                }})
                setSelectedInvoiceItemsDetails(selectedItems);
            }
        ).catch((error : any) => {
            api.error({
                message: `Create external company failed`,
                description: error.message,
                placement: 'bottom',
                duration: 1.4
            });
        })
    };

    // Buggy - has to be fixed
    const onRemoveInvoiceItem = (id : string) => {
        const newInvoiceItems = invoiceItems.filter((item : string) => item !== id);
        onChangeInvoiceItems(newInvoiceItems);
    }



    const getTotalActualPrice = selectedInvoiceItemsDetails.reduce((acc : number, item : InvoiceItem) => acc + item.actual_price, 0);
    const discountAmount = getTotalActualPrice * (discountPercentage / 100);
    const actualPrice = getTotalActualPrice - discountAmount;
    const getTotalTax = selectedInvoiceItemsDetails.reduce((acc : number, item : InvoiceItem) => acc + item.actual_tax, 0);
    const taxDiscountAmount = getTotalTax * (discountPercentage / 100);
    const actualTax = getTotalTax - taxDiscountAmount;

    const onSubmit = () => {
        createInvoice(
            companyId,
            projectId,
            subProjectId,
            invoiceDisplayName,
            invoiceDescription,
            isStatisticsInvoice,
            invoiceItems,
            actualPrice,
            getTotalActualPrice,
            discountPercentage,
            discountAmount,
            getTotalTax,
            actualTax,
            invoiceDate,
            dueDate,
            false, // paid
            status,
            paymentDate,
            loggedInUserId,
        ).then(response => {
            if (response?.error || !response?.data) {
                api.error({
                    message: `Create project invoice failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            api.info({
                message: response.message,
                placement: 'bottom',
                duration: 1.4
            });
        })
        .catch(error => {
            api.error({
                message: `Error creating project invoice`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    };

    const projectOptions = Object.keys(projects).map(projectId => ({ 
        label: projects[projectId].name, 
        value: projects[projectId].id
    }));

    const vendorOptions = Object.keys(externalCompanies).map(companyId => ({
        value: externalCompanies[companyId].id,
        label: externalCompanies[companyId].company_name
    }))


    const inviceItemsOptions = Object.keys(allInvoiceItems).map(invoiceItemId => ({
        label: getProductName(allInvoiceItems[invoiceItemId].product_id), 
        value: allInvoiceItems[invoiceItemId].id
    }));

  return (
        <Card style={{maxWidth: '1600px'}}>
            {contextHolder}
            <Row>
                <Col span={8} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Invoice name</Text>
                        <Input
                            placeholder="Invoice name"
                            onChange={(event : any) => setInvoiceDisplayName(event.target.value)}
                            onBlur={(event : any) => setInvoiceDisplayName(event.target.value)}
                            value={invoiceDisplayName}
                        />
                        <Text strong>Status</Text>
                        <Select
                            placeholder="Status"
                            style={{width: '100%'}}
                            options={statusOptions}
                            onChange={onChangeStatus}
                        />
                        <Text strong>Invoice description</Text>
                        <TextArea
                            placeholder="Invoice description"
                            onChange={(event : any) => setInvoiceDescription(event.target.value)}
                            onBlur={(event : any) => setInvoiceDescription(event.target.value)}
                            value={invoiceDescription}
                        />
                        <Text strong>Vendor</Text>
                        <Select 
                            style={{width: '100%'}}
                            value={companyId} 
                            onChange={onChangeCompany} 
                            options={vendorOptions}
                        />
                        <Text strong>Project</Text>
                        <Select
                            placeholder="Project" 
                            style={{width: '100%'}}
                            options={projectOptions}
                            onChange={onChangeProject}
                            value={projectId}
                        />
                        <Text strong>Sub project</Text>
                        <Select
                            placeholder="Sub project" 
                            style={{width: '100%'}}
                            options={subProjectOptions}
                            disabled={Boolean(!projectId)}
                            onChange={onChangeSubProject}
                            value={subProjectId}
                        />
                        <Text strong>Is this a statistical invoice?</Text>
                        <Checkbox
                            checked={isStatisticsInvoice}
                            onChange={(event : any) => setIsStatisticsInvoice(event.target.checked)}
                        />
                    </Space>
                </Col>
                <Col span={16} style={{padding: '12px'}}>
                    <Card style={{padding: '0px'}}>
                        <Space direction="vertical" style={{width: '100%', marginBottom: '16px'}}>
                    <Title level={4} style={{marginBottom: '0px'}}>Invoice items</Title>
                    <Select
                            showSearch
                            mode="multiple"
                            placeholder="Invoice items"
                            style={{width: '50%'}}
                            options={inviceItemsOptions}
                            onChange={onChangeInvoiceItems}
                            value={invoiceItems}
                        />
                        </Space>
                    {selectedInvoiceItemsDetails.length > 0 && (
                        <Table
                            columns={invoiceItemsColumns}
                            dataSource={selectedInvoiceItemsDetails}
                            pagination={false}
                        />
                    )}
                    </Card>
                    <div style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%', padding: '0px 16px'}}>
                    <div style={{ display: 'flex', justifyContent: 'flex-end', alignItems: 'center', gap: '16px', marginBottom: '16px', marginTop: '16px'}}>
                        <Space direction="vertical">
                            <Text strong>Invoice date: </Text>
                            <DatePicker 
                                onChange={onChangeInvoiceDate} 
                            />
                        </Space>
                        <Space direction="vertical">
                            <Text strong>Due date: </Text>
                            <DatePicker 
                                onChange={onChangeDueDate} 
                            />
                        </Space>
                        <Space direction="vertical">
                            <Text strong>Payment date: </Text>
                            <DatePicker
                                onChange={onChangePaymentDate}
                            />
                        </Space>   
                        </div>                     
                        <Space direction="vertical">
                            <Text strong style={{marginBottom: '80px'}}>Discount percentage</Text>
                            <Input
                                placeholder="Discount percentage"
                                type="number"
                                step="0.01"
                                max="1"
                                min="0"
                                defaultValue={0}
                                value={discountPercentage}
                                onChange={event => setDiscountPercentage(Number(event.target.value))}
                                onBlur={event => setDiscountPercentage(Number(event.target.value))}
                                suffix="%"
                            />
                        </Space>
                    </div>
                </Col>
                </Row>
            <Row>
                <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'flex-end', alignContent: 'center', gap: '24px', marginBottom: '20px', paddingRight: '40px'}}>
                    <div>
                        <Text strong>Discount</Text><br/>
                        {discountAmount} SEK<br/>
                    </div>
                    <div>
                        <Text strong>Total amount</Text><br/>
                        {formatNumberWithSpaces(actualPrice)} SEK
                    </div>
                    <div>
                        <Text strong>Total tax</Text><br/>
                        {formatNumberWithSpaces(actualTax)} SEK
                    </div>
                </div>
                </Col>
                <Col span={24}>
                    <div style={{display: 'flex', justifyContent: 'flex-end', marginRight: '16px'}}>
                        <Button type="primary" onClick={onSubmit}>Create invoice</Button>
                    </div>
                </Col>
            </Row>
        </Card>
  );
};

export default CreateInvoice;