/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { Typography, Button, Table, Card } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { useNavigate } from 'react-router-dom';
import { ExternalCompanyStatus } from '../tags/ExternalCompanyStatus';
import { getAllExternalCompanies } from '../../api/externalCompanies/getAll';
import { ExternalCompany } from '../../types/externalCompany';
import { ZoomInOutlined } from '@ant-design/icons';

const { Text } = Typography;

const economicsColumns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'ORG number',
        dataIndex: 'registration_number',
        key: 'registration_number'
    },
    {
        title: 'Contact person',
        dataIndex: 'contact_person',
        key: 'contact_person'
    },
    {
      title: 'Phone',
      dataIndex: 'phone',
      key: 'phone'
    },
    {
        title: 'Email',
        dataIndex: 'email',
        key: 'email'

    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status'
    },
    {
      title: '',
      dataIndex: 'operations',
      key: 'operations'
    },
  ];

const Expenses = ({ project } : { project: string }) => {
    const navigate = useNavigate();
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const [externalCompanies, setExternalCompanies] = useState<Array<any>>([]);

    useEffect(() => {
        if (loggedInUserId) {
            getAllExternalCompanies(loggedInUserId).then(response => {
                setExternalCompanies(response)
            }).catch(error => {
                console.log('error fetching', error)
            })
        } 
      }, [loggedInUserId]);

      const externalCompanyData: Array<any> = useMemo(() => {
        const expensesListItem = externalCompanies && externalCompanies.map((company : ExternalCompany) => {
        return {                    
            name: <Button type="link" onClick={() => navigate(`/external-company/${company.id}`)}>{company.company_name}</Button>,
            registration_number: <Text>{company.company_registration_number}</Text>,
            contact_person: <Text>{company.contact_person}</Text>,
            phone: <Text>{company.contact_phone}</Text>,
            email: <Text>{company.contact_email}</Text>,
            status: <ExternalCompanyStatus status={company.status}/>,
            operations: <Button type="link" onClick={() => navigate(`/external-company/${company.id}`)}><ZoomInOutlined /></Button>
          }
        })
        return expensesListItem;
    }, [project, externalCompanies])

    if (!externalCompanies) return null;

    return  (
        <Card 
            style={{ height: 'fit-content', padding: 0}}
            bodyStyle={{padding: '0px'}}
            >
            <Table size="small" columns={economicsColumns} dataSource={externalCompanyData} />
        </Card>
    );

}

export default Expenses;