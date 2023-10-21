/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Row, Typography } from 'antd';
import { useSelector } from 'react-redux';
import { Button, Input, Space, notification, Select } from 'antd';
import { State } from '../../types/state';
import { externalCompanyOptions } from './options';
import { updateExternalCompany } from '../../api/externalCompanies/update';
import { ExternalCompany } from '../../types/externalCompany';

const { Text } = Typography;

const UpdateProjectExpense = ({ externalCompany, setEditing } : { externalCompany : ExternalCompany, setEditing : (open : boolean) => void }) => {
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const allProjects = useSelector((state: State) => state.application.projects);
    const [companyName, setCompanyName] = useState('');
    const [companyRegistrationNumber, setCompanyRegistrationNumber] = useState('');
    const [contactPerson, setContactPerson] = useState('');
    const [contactEmail, setContactEmail] = useState('');
    const [contactPhone, setContactPhone] = useState('');
    const [address, setAddress] = useState('');
    const [city, setCity] = useState('');
    const [stateProvince, setStateProvince] = useState('');
    const [country, setCountry] = useState('');
    const [postalCode, setPostalCode] = useState('');
    const [paymentTerms, setPaymentTerms] = useState('');
    const [billingCurrency, setBillingCurrency] = useState('');
    const [bankAccountInfo, setBankAccountInfo] = useState('');
    const [taxIdentificationNumber, setTaxIdentificationNumber] = useState('');
    const [status, setStatus] = useState('');
    const [assignedProjects, setAssignedProjects] = useState<Array<string>>([]);
    const [invoicePending, setInvoicePending] = useState<Array<string>>([]);
    const [invoiceHistory, setInvoiceHistory] = useState<Array<string>>([]);
    const [contractualAgreements, setContractualAgreements] = useState<Array<string>>([]);

    useEffect(() => {
        setCompanyName(externalCompany.company_name);
        setCompanyRegistrationNumber(externalCompany.company_registration_number);
        setContactPerson(externalCompany.contact_person);
        setContactEmail(externalCompany.contact_email);
        setContactPhone(externalCompany.contact_phone);
        setAddress(externalCompany.address);
        setCity(externalCompany.city);
        setStateProvince(externalCompany.state_province);
        setCountry(externalCompany.country);
        setPostalCode(externalCompany.postal_code);
        setPaymentTerms(externalCompany.payment_terms);
        setBillingCurrency(externalCompany.billing_currency);
        setBankAccountInfo(externalCompany.bank_account_info);
        setTaxIdentificationNumber(externalCompany.tax_identification_number);
        setStatus(externalCompany.status);
        setAssignedProjects(externalCompany.assigned_projects);
        setInvoicePending(externalCompany.invoice_pending);
        setInvoiceHistory(externalCompany.invoice_history);
        setContractualAgreements(externalCompany.contractual_agreements);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const projectOptions = allProjects.map(project => {
        return { label: project.name, value: project.id}
      }
    );

    const onChangeStatus = (value : any) => setStatus(value);
    const onChangeAssignedProjects = (value: any) => setAssignedProjects(value);
    const onChangePendingInvoices = (value : any) => setInvoicePending(value); 
    const onChangeInvoiceHistory = (value : any) => setInvoiceHistory(value);
    const onChangeContractualAgreements = (value : any) => setContractualAgreements(value);


    const onSubmit = () => {
        updateExternalCompany(
            userId,
            externalCompany.id,
            companyName,
            companyRegistrationNumber,
            contactPerson,
            contactEmail,
            contactPhone,
            address,
            city,
            stateProvince,
            country,
            postalCode,
            paymentTerms,
            billingCurrency,
            bankAccountInfo,
            taxIdentificationNumber,
            status,
            assignedProjects,
            invoicePending,
            invoiceHistory,
            contractualAgreements,
        ).then(response => {
            if (response?.error) {
                api.error({
                    message: `Create external company failed`,
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
                message: `Error creating external company`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    };

  return (
        <>
            {contextHolder}
            <Row>
                <Col span={24}>
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px'}}>
                        <Button onClick={() => setEditing(false)}>Close</Button>
                        <Button type="primary" onClick={onSubmit}>Save</Button>
                    </div>
                </Col>
            </Row>
            <Row>
                <Col span={12} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Company name</Text>
                        <Input
                            placeholder="Project" 
                            style={{width: '100%'}}
                            onChange={event => setCompanyName(event.target.value)}
                            value={companyName}
                        />
                        <Text strong>Company registration number</Text>
                        <Input
                            placeholder="Company registration number" 
                            style={{width: '100%'}}
                            onChange={event => setCompanyRegistrationNumber(event.target.value)}
                            value={companyRegistrationNumber}
                        />
                        <Text strong>Contact person</Text>
                        <Input
                            placeholder="Contact person" 
                            style={{width: '100%'}}
                            onChange={event => setContactPerson(event.target.value)}
                            value={contactPerson}
                        />
                        <Text strong>Contact email</Text>
                        <Input
                            placeholder="Contact email" 
                            style={{width: '100%'}}
                            onChange={event => setContactEmail(event.target.value)}
                            value={contactEmail}
                        />
                        <Text strong>Contact phone</Text>
                        <Input
                            placeholder="Contact phone" 
                            style={{width: '100%'}}
                            onChange={event => setContactPhone(event.target.value)}
                            value={contactPhone}
                        />
                        </Space>
                        </Col>
                        <Col span={12} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Address</Text>
                        <Input
                            placeholder="Address" 
                            style={{width: '100%'}}
                            onChange={event => setAddress(event.target.value)}
                            value={address}
                        />
                        <Text strong>City</Text>
                        <Input
                            placeholder="City" 
                            style={{width: '100%'}}
                            onChange={event => setCity(event.target.value)}
                            value={city}
                        />
                        <Text strong>State province</Text>
                        <Input
                            placeholder="State province" 
                            style={{width: '100%'}}
                            onChange={event => setStateProvince(event.target.value)}
                            value={stateProvince}
                        />
                        <Text strong>Country</Text>
                        <Input
                            placeholder="Country" 
                            style={{width: '100%'}}
                            onChange={event => setCountry(event.target.value)}
                            value={country}
                        />
                        <Text strong>Postal code</Text>
                        <Input
                            placeholder="Postal code" 
                            style={{width: '100%'}}
                            onChange={event => setPostalCode(event.target.value)}
                            value={postalCode}
                        />
                        </Space>
                        </Col>
                        </Row>
                        <Row>
                        <Col span={12} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Payment terms</Text>
                        <Input
                            placeholder="Payment terms" 
                            style={{width: '100%'}}
                            onChange={event => setPaymentTerms(event.target.value)}
                            value={paymentTerms}
                        />
                        <Text strong>Billing currency</Text>
                        <Input
                            placeholder="Billing currency" 
                            style={{width: '100%'}}
                            onChange={event => setBillingCurrency(event.target.value)}
                            value={billingCurrency}
                        />
                        <Text strong>Bank account info</Text>
                        <Input
                            placeholder="Bank account info" 
                            style={{width: '100%'}}
                            onChange={event => setBankAccountInfo(event.target.value)}
                            value={bankAccountInfo}
                        />
                        <Text strong>Tax identification number</Text>
                         <Input
                            placeholder="Tax identification number" 
                            style={{width: '100%'}}
                            onChange={event => setTaxIdentificationNumber(event.target.value)}
                            value={taxIdentificationNumber}
                        />
                        </Space>
                        </Col>
                        <Col span={12} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Status</Text>
                        <Select
                            placeholder="Select status"
                            style={{width: '100%'}}
                            options={externalCompanyOptions}
                            onChange={onChangeStatus}
                            value={status}
                        />
                        <Text strong>Assigned projects</Text>
                        <Select
                            mode="multiple"
                            style={{width: '100%'}}
                            options={projectOptions}
                            onChange={onChangeAssignedProjects}
                            value={assignedProjects}
                        />      
                        <Text strong>Pending invoices</Text>
                        <Select
                            mode="multiple"
                            style={{width: '100%'}}
                            options={[
                                {value: 'invoice-one', label: 'invoice-one'},
                                {value: 'invoice-two', label: 'invoice-two'},
                                {value: 'invoice-three', label: 'invoice-three'},
                            ]}
                            onChange={onChangePendingInvoices}
                            value={invoicePending}
                        />
                        <Text strong>Invoice history</Text>
                        <Select
                            mode="multiple"
                            style={{width: '100%'}}
                            options={[
                                {value: 'invoice-one-history', label: 'invoice-one-history'},
                                {value: 'invoice-two-history', label: 'invoice-two-history'},
                                {value: 'invoice-three-history', label: 'invoice-three-history'},
                            ]}
                            onChange={onChangeInvoiceHistory}
                            value={invoiceHistory}
                        />
                        <Text strong>Contractual agreements</Text>
                        <Select
                            mode="multiple"
                            style={{width: '100%'}}
                            options={[
                                {value: 'contractualAgreements-one', label: 'contractualAgreements-one'},
                                {value: 'contractualAgreements-two', label: 'contractualAgreements-two'},
                                {value: 'contractualAgreements-three', label: 'contractualAgreements-three'},
                            ]}
                            onChange={onChangeContractualAgreements}
                            value={contractualAgreements}
                        />
                    </Space>
                </Col>
            </Row>
            <Row>
                <Col span={24}>
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px'}}>
                        <Button onClick={() => setEditing(false)}>Close</Button>
                        <Button type="primary" onClick={onSubmit}>Save</Button>
                    </div>
                </Col>
            </Row>
        </>
  );
};

export default UpdateProjectExpense;
