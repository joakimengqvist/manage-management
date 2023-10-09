/*


type ExternalCompany struct {
	CompanyName               string    `json:"company_name" sql:"not null"`
	CompanyRegistrationNumber string    `json:"company_registration_number"`
	ContactPerson             string    `json:"contact_person"`
	ContactEmail              string    `json:"contact_email"`
	ContactPhone              string    `json:"contact_phone"`
	Address                   string    `json:"address"`
	City                      string    `json:"city"`
	StateProvince             string    `json:"state_province"`
	Country                   string    `json:"country"`
	PostalCode                string    `json:"postal_code"`
	PaymentTerms              string    `json:"payment_terms"`
	BillingCurrency           string    `json:"billing_currency"`
	BankAccountInfo           string    `json:"bank_account_info"`
	TaxIdentificationNumber   string    `json:"tax_identification_number"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Status                    string    `json:"status"`
	AssignedProjects          []string  `json:"assigned_projects"`
	InvoicePending            []string  `json:"invoice_pending"`
	InvoiceHistory            []string  `json:"invoice_history"`
	ContractualAgreements     []string  `json:"contractual_agreements"`
}

*/

/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Col, Row, Typography } from 'antd';
import { useSelector } from 'react-redux';
import { Button, Input, Space, Card, notification, Select } from 'antd';
import { State } from '../../types/state';
import { cardShadow } from '../../enums/styles';
import { createExternalCompany } from '../../api/externalCompanies/create';
import { externalCompanyOptions } from './options';

const { Title, Text } = Typography;

// const numberPattern = /^[0-9]+$/;

const CreateProjectExpense: React.FC = () => {
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
        createExternalCompany(
            userId,
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
            if (response?.error || !response?.data) {
                api.error({
                    message: `Create external company failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            api.info({
                message: `Created external company`,
                description: 'Succesfully created external company.',
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

    const setMockedData = () => {
        setCompanyName('Engqvist staff')
        setCompanyRegistrationNumber(generateRandomNumberString());
        setContactPerson('Joakim Engqvist');
        setContactEmail('joakim@engqvist.se');
        setContactPhone('0730522473');
        setAddress('Luxgatan 8');
        setCity('Stockholm');
        setStateProvince('Stockholm');
        setCountry('Sweden');
        setPostalCode('11262');
        setPaymentTerms('payment-terms-partner');
        setBillingCurrency('SEK');
        setBankAccountInfo(generateRandomStringIBAN());
        setTaxIdentificationNumber(generateRandomNumberString());
        setStatus('active');
        setInvoicePending(['invoice-one', 'invoice-two'])
        setInvoiceHistory(['invoice-one-history', 'invoice-two-history'])
        setContractualAgreements(['contractualAgreements-one', 'contractualAgreements-two'])
    }

  return (
        <Card bordered={false} style={{borderRadius: 0, boxShadow: cardShadow, maxWidth: '1100px'}}>
            {contextHolder}
            <Title level={4}>Create external company</Title>
            <Row>
                <Col span={6} style={{padding: '12px 12px 12px 0px'}}>
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
                        <Col span={6} style={{padding: '12px 12px 12px 0px'}}>
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
                        <Col span={6} style={{padding: '12px 12px 12px 0px'}}>
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
                        <Col span={6} style={{padding: '12px 12px 12px 0px'}}>
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
                <Col>
                    <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                        <Button onClick={setMockedData}>Populate with mock</Button>
                        <Button type="primary" onClick={onSubmit}>Create external company</Button>
                    </div>
                </Col>
            </Row>
        </Card>
  );
};

export default CreateProjectExpense;

function generateRandomNumberString() {
    let result = '';
    const numbers = '0123456789';
  
    for (let i = 0; i < 10; i++) {
      const randomIndex = Math.floor(Math.random() * numbers.length);
      result += numbers.charAt(randomIndex);
    }
  
    return result;
  }

  function generateRandomStringIBAN() {
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
  
    for (let i = 0; i < 20; i++) {
      const randomIndex = Math.floor(Math.random() * characters.length);
      result += characters.charAt(randomIndex);
    }
  
    return result;
  }