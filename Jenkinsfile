pipeline {
    agent {
        docker {
            image 'node:23'
            args '--user root'
        }
    }
    environment {
        front = 'https://github.com/skpub/IMUweb'
        article = 'https://github.com/skpub/IMUarticle'
        private_key = credentials('SSH_SECRET_KEY_IMUWEB')
        deploy_host = credentials('IMU_DEPLOY_HOST')
        deploy_user = credentials('IMU_DEPLOY_USER')
        deploy_path = credentials('IMU_DEPLOY_PATH')
    }
    stages {
        stage('Prepare') {
            steps {
                echo 'Installing rsync..'
                sh 'apt-get update && apt-get install -y rsync'
            }
        }
        stage('Checkout Repositories') {
            steps {
                echo 'Checking out repositories..'
                dir('IMUweb') {
                    git url: "${front}", branch: 'master'
                    dir ('IMUarticle') {
                        git url: "${article}", branch: 'master'
                    }
                }
            }
        }
        stage('Install Dependencies') {
            steps {
                echo 'Installing dependencies..'
                dir('IMUweb') {
                    sh 'npm install'
                }
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
                dir('IMUweb') {
                    sh 'npm run build'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                dir('IMUweb') {
                    withCredentials([sshUserPrivateKey(credentialsId: "SSH_SECRET_KEY_IMUWEB", keyFileVariable: 'SSH_KEY')]) {
                        sh '''
                            chmod 600 "${SSH_KEY}"
                            export RSYNC_RSH="ssh -i ${SSH_KEY} -o StrictHostKeyChecking=no"
                            rsync -avz ./build/ ${deploy_user}@${deploy_host}:${deploy_path}
                        '''
                    }
                }
            }
        }
    }
}
