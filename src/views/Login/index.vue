<template>
    <div class="login-container">
        <el-form
            :model="ruleForm"
            status-icon
            ref="ruleForm"
            label-width="100px"
            class="login-ruleForm"
        >
            <h1 class="form-title">系统登陆</h1>
            <el-form-item label="账号：" prop="pass">
                <el-input v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码：" prop="checkPass">
                <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
export default {
    data() {
        return {
            ruleForm: {
                pass: '',
                checkPass: ''
            }
        };
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    if(this.ruleForm.pass === "admin" && this.ruleForm.checkPass === "88888888"){
                        localStorage.setItem("userToken",JSON.stringify(this.ruleForm))
                        this.$router.push("/list")
                        this.$message.success('登陆成功！'); 
                    }else{
                       this.$message.error('密码错误！'); 
                    }
                } else {
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        }
    }
}
</script>
<style>
.login-container {
    justify-content: center;
    height: 100%;
    display: flex;
    align-items: center;
}

.login-ruleForm {
    display: inline-block;
    width: 500px;
}

.form-title {
    margin-bottom: 20px;
    text-align: center;
}
</style>