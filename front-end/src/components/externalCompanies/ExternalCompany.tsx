/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Space, Typography, Row, Col } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { cardShadow } from '../../enums/styles';
import { getExternalCompanyById } from '../../api/externalCompanies/getById';
import { ExternalCompany } from '../../types/externalCompany';

const { Text, Title } = Typography;

const ExternalCompanyDetails: React.FC = () => {
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const [externalCompany, setExternalCompany] = useState<null | ExternalCompany>(null);
    const { id } =  useParams(); 
    const externalCompanyId = id || '';

    useEffect(() => {
        if (loggedInUserId) {
            getExternalCompanyById(loggedInUserId, externalCompanyId).then(response => {
                setExternalCompany(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId]);

    return (
        <Card bordered={false} style={{ borderRadius: 0, boxShadow: cardShadow}}>
            <Title level={4}>External company information</Title>
                {externalCompany && (
                    <Row>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Name</Text>
                                {externalCompany.company_name}
                                <Text strong>Counter</Text>
                                {externalCompany.country}
                                <Text strong>City</Text>
                                {externalCompany.city}
                                <Text strong>State / Province</Text>
                                {externalCompany.state_province}
                                <Text strong>Address</Text>
                                {externalCompany.address}
                                <Text strong>Postal code</Text>
                                {externalCompany.postal_code}
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Contact person</Text>
                                {externalCompany.contact_person}
                                <Text strong>email</Text>
                                {externalCompany.contact_email}
                                <Text strong>phone</Text>
                                {externalCompany.contact_phone}
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Registration number</Text>
                                {externalCompany.company_registration_number}
                                <Text strong>Tax identification number</Text>
                                {externalCompany.tax_identification_number}
                                <Text strong>Billing currency</Text>
                                {externalCompany.billing_currency}
                                <Text strong>Bank account info</Text>
                                {externalCompany.bank_account_info}
                                <Text strong>ID</Text>
                                {externalCompany.id}
                            </Space>
                        </Col>
                    </Row>
                )}
        </Card>
    )

}

export default ExternalCompanyDetails;