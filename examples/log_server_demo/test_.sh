config_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${config_dir}

${config_dir}/build_linux.sh

scp log_server_demo_linux ecs-user@182.92.99.61:
scp .env ecs-user@182.92.99.61:

ssh ecs-user@182.92.99.61 '
    if [ ! -d 'test1' ]; then mkdir test1; fi;
    mv log_server_demo_linux test1
    mv .env test1
'